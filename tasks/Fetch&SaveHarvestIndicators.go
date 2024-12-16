package tasks

import (
	"encoding/json"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

const HarvestIndicatorsAPIURL = "https://operations.cropwise.com/api/v3/harvest_indicators"

type HarvestIndicatorsResponse struct {
	Data []models.HarvestIndicatorModel `json:"data"`
	Meta struct {
		Response struct {
			ObtainedRecords int `json:"obtained_records"`
			LastRecordID    int `json:"last_record_id"`
		} `json:"response"`
	} `json:"meta"`
}

// FetchAndSaveHarvestIndicators - загрузка данных и сохранение в базу данных
func FetchAndSaveHarvestIndicators(token string, schemaName string) error {
	log.Printf("Начинаем загрузку данных Harvest Indicators для схемы: %s", schemaName)

	// Устанавливаем search_path для схемы компании
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка установки search_path на %s: %v", schemaName, err)
		return err
	}
	defer func() {
		resetSearchPath := "SET search_path TO public"
		if err := database.DB.Exec(resetSearchPath).Error; err != nil {
			log.Printf("Ошибка сброса search_path на public: %v", err)
		}
	}()

	client := &http.Client{}
	fromID := 0

	for {
		url := HarvestIndicatorsAPIURL + "?from_id=" + strconv.Itoa(fromID)
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

		var response HarvestIndicatorsResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}

		// Сохраняем данные в базу
		for _, indicator := range response.Data {
			if err := database.DB.Save(&indicator).Error; err != nil {
				return fmt.Errorf("ошибка сохранения Harvest Indicator с ID %d: %w", indicator.ID, err)
			}
		}

		// Проверяем, нужно ли продолжать загружать данные
		if response.Meta.Response.ObtainedRecords == 0 {
			break
		}

		fromID = response.Meta.Response.LastRecordID + 1
	}

	return nil
}
