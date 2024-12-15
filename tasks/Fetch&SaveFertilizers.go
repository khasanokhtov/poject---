package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FertilizersAPIURL = "https://operations.cropwise.com/api/v3/fertilizers"

type FertilizersResponse struct {
	Data []models.FertilizerModel `json:"data"`
}

func FetchAndSaveFertilizers(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FertilizersAPIURL, nil)
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
		return errors.New("не удалось получить данные о удобрениях")
	}

	// Декодируем ответ
	var fertilizersResponse FertilizersResponse
	if err := json.NewDecoder(resp.Body).Decode(&fertilizersResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, fertilizer := range fertilizersResponse.Data {
		if err := database.DB.Create(&fertilizer).Error; err != nil {
			return fmt.Errorf("ошибка сохранения удобрения с ID %d: %w", fertilizer.ID, err)
		}
	}

	return nil
}
