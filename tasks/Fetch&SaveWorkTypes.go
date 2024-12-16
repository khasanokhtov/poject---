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

// Константа для URL API
const WorkTypesAPIURL = "https://operations.cropwise.com/api/v3/work_types"

// FetchAndSaveWorkTypes - функция для загрузки и сохранения данных work_types
func FetchAndSaveWorkTypes(token, schemaName string) error {
	log.Printf("Начинаем загрузку данных work_types для схемы: %s", schemaName)

	// Установить search_path для работы с нужной схемой
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

	// Цикл для обработки пагинации
	for {
		url := WorkTypesAPIURL + "?from_id=" + strconv.Itoa(fromID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("ошибка создания запроса: %w", err)
		}
		req.Header.Set("X-User-Api-Token", token)

		// Выполнить запрос
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("ошибка выполнения запроса: %w", err)
		}
		defer resp.Body.Close()

		// Проверить статус ответа
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("ошибка ответа от API: %s, тело: %s", resp.Status, string(body))
		}

		// Прочитать тело ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения тела ответа: %w", err)
		}

		// Парсинг JSON
		var response struct {
			Data []models.WorkType `json:"data"`
			Meta struct {
				Response struct {
					ObtainedRecords int `json:"obtained_records"`
					LastRecordID    int `json:"last_record_id"`
				} `json:"response"`
			} `json:"meta"`
		}
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}

		// Сохранение данных в базу
		for _, workType := range response.Data {
			if err := database.DB.Save(&workType).Error; err != nil {
				log.Printf("Ошибка сохранения work_type с ID %d: %v", workType.ID, err)
			}
		}

		// Проверка завершения загрузки данных
		if response.Meta.Response.ObtainedRecords == 0 {
			break
		}

		// Обновить fromID для следующего запроса
		fromID = response.Meta.Response.LastRecordID + 1
	}

	log.Printf("Данные work_types успешно загружены для схемы %s", schemaName)
	return nil
}
