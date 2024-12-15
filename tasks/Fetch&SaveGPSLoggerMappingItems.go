package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const GPSLoggerMappingItemsAPIURL = "https://operations.cropwise.com/api/v3/gps_logger_mapping_items"

type GPSLoggerMappingItemsResponse struct {
	Data []models.GPSLoggerMappingItemModel `json:"data"`
}

// FetchAndSaveGPSLoggerMappingItems - загрузка данных и сохранение в базу данных
func FetchAndSaveGPSLoggerMappingItems(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", GPSLoggerMappingItemsAPIURL, nil)
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}
	req.Header.Set("X-User-Api-Token", token)

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return errors.New("не удалось получить данные о GPS Logger Mapping Items")
	}

	// Декодируем ответ
	var gpsLoggerMappingItemsResponse GPSLoggerMappingItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&gpsLoggerMappingItemsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, mappingItem := range gpsLoggerMappingItemsResponse.Data {
		if err := database.DB.Create(&mappingItem).Error; err != nil {
			return fmt.Errorf("ошибка сохранения GPS Logger Mapping Item с ID %d: %w", mappingItem.ID, err)
		}
	}

	return nil
}
