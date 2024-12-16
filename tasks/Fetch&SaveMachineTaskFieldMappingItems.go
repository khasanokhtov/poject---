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

// Константа для эндпоинта
const MachineTaskFieldMappingItemsAPIURL = "https://operations.cropwise.com/api/v3/machine_task_field_mapping_items"

// FetchAndSaveMachineTaskFieldMappingItems - функция для загрузки и сохранения данных
func FetchAndSaveMachineTaskFieldMappingItems(token, schemaName string) error {
	log.Printf("Начинаем загрузку данных для схемы: %s", schemaName)

	// Устанавливаем search_path для работы с нужной схемой
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

	// Цикл для пагинации
	for {
		url := MachineTaskFieldMappingItemsAPIURL + "?from_id=" + strconv.Itoa(fromID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("ошибка создания запроса: %w", err)
		}
		req.Header.Set("X-User-Api-Token", token)

		// Выполняем запрос
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("ошибка выполнения запроса: %w", err)
		}
		defer resp.Body.Close()

		// Проверяем статус ответа
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("ошибка ответа: %s", resp.Status)
		}

		// Читаем тело ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %w", err)
		}

		// Парсим JSON-ответ
		var response struct {
			Data []models.MachineTaskFieldMappingItemModel `json:"data"`
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

		// Сохраняем данные в базу
		for _, item := range response.Data {
			if err := database.DB.Save(&item).Error; err != nil {
				return fmt.Errorf("ошибка сохранения machine_task_field_mapping_item с ID %d: %w", item.ID, err)
			}
		}

		// Проверяем на окончание данных
		if response.Meta.Response.ObtainedRecords == 0 {
			break
		}

		// Обновляем fromID для следующего запроса
		fromID = response.Meta.Response.LastRecordID + 1
	}

	log.Printf("Данные machine_task_field_mapping_items успешно загружены для схемы %s", schemaName)
	return nil
}
