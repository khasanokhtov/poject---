package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FieldScoutReportsAPIURL = "https://operations.cropwise.com/api/v3/field_scout_reports"

type FieldScoutReportsResponse struct {
	Data []models.FieldScoutReportModel `json:"data"`
}

func FetchAndSaveFieldScoutReports(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FieldScoutReportsAPIURL, nil)
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
		return errors.New("не удалось получить данные о скаут-отчетах")
	}

	// Декодируем ответ
	var fieldScoutReportsResponse FieldScoutReportsResponse
	if err := json.NewDecoder(resp.Body).Decode(&fieldScoutReportsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, fieldScoutReport := range fieldScoutReportsResponse.Data {
		if err := database.DB.Create(&fieldScoutReport).Error; err != nil {
			return fmt.Errorf("ошибка сохранения скаут-отчета с ID %d: %w", fieldScoutReport.ID, err)
		}
	}

	return nil
}
