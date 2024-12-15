package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FuelHourlyDataItemsAPIURL = "https://operations.cropwise.com/api/v3/fuel_hourly_data_items"

type FuelHourlyDataItemsResponse struct {
	Data []models.FuelHourlyDataItemModel `json:"data"`
}

func FetchAndSaveFuelHourlyDataItems(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FuelHourlyDataItemsAPIURL, nil)
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
		return errors.New("не удалось получить данные о почасовом расходе топлива")
	}

	// Декодируем ответ
	var fuelHourlyDataItemsResponse FuelHourlyDataItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&fuelHourlyDataItemsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, fuelHourlyDataItem := range fuelHourlyDataItemsResponse.Data {
		if err := database.DB.Create(&fuelHourlyDataItem).Error; err != nil {
			return fmt.Errorf("ошибка сохранения записи с ID %d: %w", fuelHourlyDataItem.ID, err)
		}
	}

	return nil
}
