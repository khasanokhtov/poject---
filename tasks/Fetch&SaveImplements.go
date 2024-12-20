package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"integration-cropwise-v1/models"

	"gorm.io/gorm"
)

const ImplementsAPIURL = "https://operations.cropwise.com/api/v3/implements"

type ImplementsResponse struct {
	Data []models.Implement `json:"data"`
	Meta struct {
		Response struct {
			ObtainedRecords int `json:"obtained_records"`
			LastRecordID    int `json:"last_record_id"`
		} `json:"response"`
	} `json:"meta"`
}

func FetchAndSaveImplements(db *gorm.DB, token, schemaName string) error {
	log.Printf("Начинаем загрузку данных Implements для схемы: %s", schemaName)

	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := db.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка установки search_path на %s: %v", schemaName, err)
		return err
	}
	defer func() {
		resetSearchPath := "SET search_path TO public"
		if err := db.Exec(resetSearchPath).Error; err != nil {
			log.Printf("Ошибка сброса search_path на public: %v", err)
		}
	}()

	client := &http.Client{}
	fromID := 0

	for {
		url := ImplementsAPIURL + "?from_id=" + strconv.Itoa(fromID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("ошибка создания запроса: %w", err)
		}
		req.Header.Set("X-User-Api-Token", token)

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("ошибка выполнения запроса: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("ошибка ответа: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %w", err)
		}

		var implementsResponse ImplementsResponse
		if err := json.Unmarshal(body, &implementsResponse); err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}

		for _, implement := range implementsResponse.Data {
			if err := db.Save(&implement).Error; err != nil {
				return fmt.Errorf("ошибка сохранения implement с ID %d: %w", implement.ID, err)
			}
		}

		if implementsResponse.Meta.Response.ObtainedRecords == 0 {
			break
		}

		fromID = implementsResponse.Meta.Response.LastRecordID + 1
	}

	return nil
}
