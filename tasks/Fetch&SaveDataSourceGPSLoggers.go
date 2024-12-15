package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const DataSourceGPSLoggersAPIURL = "https://operations.cropwise.com/api/v3/data_source_gps_loggers"

type GPSLoggerResponse struct {
	Data []models.DataSourceGPSLoggerModel `json:"data"`
}

func FetchAndSaveDataSourceGPSLoggers(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", GPSLoggersAPIURL, nil)
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
		return errors.New("не удалось получить данные о GPS-логгерах")
	}

	// Декодируем ответ
	var gpsLoggerResponse GPSLoggerResponse
	if err := json.NewDecoder(resp.Body).Decode(&gpsLoggerResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, gpsLogger := range gpsLoggerResponse.Data {
		if err := database.DB.Create(&gpsLogger).Error; err != nil {
			return fmt.Errorf("ошибка сохранения GPS-логгера с ID %d: %w", gpsLogger.ID, err)
		}
	}

	return nil
}
