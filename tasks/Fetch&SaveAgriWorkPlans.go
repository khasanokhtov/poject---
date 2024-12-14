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

const AgriWorkPlansEndpoint = "https://operations.cropwise.com/api/v3/agri_work_plans"

func FetchAndSaveAgriWorkPlans(token, schemaName string) error {
	log.Printf("Начинаем загрузку данных для схемы: %s", schemaName)

	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка установки search_path на %s: %v", schemaName, err)
		return err
	}
	defer resetSearchPath()

	client := &http.Client{}
	fromID := 0

	for {
		url := AgriWorkPlansEndpoint + "?from_id=" + strconv.Itoa(fromID)
		log.Printf("Запрос данных с URL: %s", url)

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
		log.Printf("Ответ успешно прочитан для схемы %s, длина ответа: %d байт", schemaName, len(body))

		var response struct {
			Data []models.AgriWorkPlanModel `json:"data"`
			Meta struct {
				Response struct {
					ObtainedRecords int `json:"obtained_records"`
					LastRecordID    int `json:"last_record_id"`
				} `json:"response"`
			} `json:"meta"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}
		log.Printf("Получено %d записей для схемы %s", len(response.Data), schemaName)

		for _, plan := range response.Data {
			if err := database.DB.Save(&plan).Error; err != nil {
				return fmt.Errorf("ошибка сохранения плана ID %d в схеме %s: %w", plan.ID, schemaName, err)
			}
		}

		if response.Meta.Response.ObtainedRecords == 0 {
			log.Printf("Все данные загружены для схемы %s", schemaName)
			break
		}

		fromID = response.Meta.Response.LastRecordID + 1
	}

	return nil
}
