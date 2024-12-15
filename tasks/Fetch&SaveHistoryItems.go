package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const HistoryItemsAPIURL = "https://operations.cropwise.com/api/v3/history_items"

// HistoryItemsResponse - структура для декодирования ответа API
type HistoryItemsResponse struct {
	Data []models.HistoryItemModel `json:"data"`
}

// FetchAndSaveHistoryItems - функция для загрузки и сохранения данных history_items
func FetchAndSaveHistoryItems(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем HTTP-запрос к внешнему API
	req, err := http.NewRequest("GET", HistoryItemsAPIURL, nil)
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
		return errors.New("не удалось получить данные о history_items")
	}

	// Декодируем JSON-ответ
	var historyItemsResponse HistoryItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&historyItemsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу данных
	for _, item := range historyItemsResponse.Data {
		if err := database.DB.Create(&item).Error; err != nil {
			return fmt.Errorf("ошибка сохранения history_item с ID %d: %w", item.ID, err)
		}
	}

	return nil
}
