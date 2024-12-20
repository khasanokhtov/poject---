package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"integration-cropwise-v1/models"

	"gorm.io/gorm"
)

const AgroOperationsEndpoint = "https://operations.cropwise.com/api/v3a/agro_operations"

func FetchAndSaveAgroOperations(db *gorm.DB, token, schemaName string) error {
	log.Printf("Начинаем загрузку агроопераций для схемы: %s", schemaName)

	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := db.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка установки search_path для %s: %v", schemaName, err)
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
		url := fmt.Sprintf("%s?from_id=%d", AgroOperationsEndpoint, fromID)
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
			return fmt.Errorf("ошибка ответа API: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения тела ответа: %w", err)
		}

		var response struct {
			Data []models.AgroOperation `json:"data"`
			Meta struct {
				Response struct {
					ObtainedRecords int `json:"obtained_records"`
					LastRecordID    int `json:"last_record_id"`
				} `json:"response"`
			} `json:"meta"`
		}
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("ошибка декодирования JSON: %w", err)
		}

		for _, operation := range response.Data {
			if err := db.Save(&operation).Error; err != nil {
				return fmt.Errorf("ошибка сохранения операции ID %d: %w", operation.ID, err)
			}
		}

		if response.Meta.Response.ObtainedRecords == 0 {
			log.Printf("Загрузка завершена для схемы: %s", schemaName)
			break
		}

		fromID = response.Meta.Response.LastRecordID + 1
		log.Printf("Загружено %d записей. Переход к from_id=%d", response.Meta.Response.ObtainedRecords, fromID)
	}

	return nil
}