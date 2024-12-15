package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const GPSLoggersAPIURL = "https://operations.cropwise.com/api/v3/gps_loggers"

type GPSLoggersResponse struct {
	Data []models.GPSLoggerModel `json:"data"`
}

// FetchAndSaveGPSLoggers - функция для загрузки данных GPS-логгеров и сохранения их в базу данных
func FetchAndSaveGPSLoggers(token string, schemaName string) error {
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
		return errors.New("не удалось получить данные GPS-логгеров")
	}

	// Декодируем ответ
	var gpsLoggersResponse GPSLoggersResponse
	if err := json.NewDecoder(resp.Body).Decode(&gpsLoggersResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, logger := range gpsLoggersResponse.Data {
		if err := database.DB.Create(&logger).Error; err != nil {
			return fmt.Errorf("ошибка сохранения GPS-логгера с ID %d: %w", logger.ID, err)
		}
	}

	return nil
}
