package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const HarvestIndicatorsAPIURL = "https://operations.cropwise.com/api/v3/harvest_indicators"

type HarvestIndicatorsResponse struct {
	Data []models.HarvestIndicatorModel `json:"data"`
}

// FetchAndSaveHarvestIndicators - загрузка данных и сохранение в базу данных
func FetchAndSaveHarvestIndicators(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", HarvestIndicatorsAPIURL, nil)
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
		return errors.New("не удалось получить данные о Harvest Indicators")
	}

	// Декодируем ответ
	var harvestIndicatorsResponse HarvestIndicatorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&harvestIndicatorsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, indicator := range harvestIndicatorsResponse.Data {
		if err := database.DB.Create(&indicator).Error; err != nil {
			return fmt.Errorf("ошибка сохранения Harvest Indicator с ID %d: %w", indicator.ID, err)
		}
	}

	return nil
}
