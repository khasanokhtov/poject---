package tasks

import (
	"encoding/json"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"io"
	"log"
	"net/http"
)

const GPSLoggersAPIURL = "https://operations.cropwise.com/api/v3/gps_loggers"

func FetchAndSaveGPSLoggers(token string, schemaName string) error {
	log.Printf("Начинаем загрузку данных GPS-логгеров для схемы: %s", schemaName)

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

	for {
		// Формируем URL с параметром from_id
		url := fmt.Sprintf("%s?from_id=%d", GPSLoggersAPIURL, fromID)
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
			return fmt.Errorf("ошибка ответа API: %s", resp.Status)
		}

		// Читаем тело ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения тела ответа: %w", err)
		}

		// Парсим JSON-ответ
		var response struct {
			Data []models.GPSLoggerModel `json:"data"`
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

		// Сохраняем данные GPS-логгеров в базу
		for _, logger := range response.Data {
			if err := database.DB.Create(&logger).Error; err != nil {
				return fmt.Errorf("ошибка сохранения GPS-логгера с ID %d: %w", logger.ID, err)
			}
		}

		// Проверяем, нужно ли продолжать загрузку
		if response.Meta.Response.ObtainedRecords == 0 {
			log.Printf("Загрузка GPS-логгеров завершена для схемы: %s", schemaName)
			break
		}

		// Обновляем fromID для следующего запроса
		fromID = response.Meta.Response.LastRecordID + 1
		log.Printf("Загружено %d записей. Переход к from_id=%d", response.Meta.Response.ObtainedRecords, fromID)
	}

	return nil
}
