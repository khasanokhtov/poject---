package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const HarvestWeighingsAPIURL = "https://operations.cropwise.com/api/v3/harvest_weighings"

type HarvestWeighingsResponse struct {
	Data []models.HarvestWeighingModel `json:"data"`
}

// FetchAndSaveHarvestWeighings - загрузка данных и сохранение в базу данных
func FetchAndSaveHarvestWeighings(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", HarvestWeighingsAPIURL, nil)
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
		return errors.New("не удалось получить данные о Harvest Weighings")
	}

	// Декодируем ответ
	var harvestWeighingsResponse HarvestWeighingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&harvestWeighingsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, weighing := range harvestWeighingsResponse.Data {
		if err := database.DB.Create(&weighing).Error; err != nil {
			return fmt.Errorf("ошибка сохранения Harvest Weighing с ID %d: %w", weighing.ID, err)
		}
	}

	return nil
}
