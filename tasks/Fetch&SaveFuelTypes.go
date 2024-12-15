package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FuelTypesAPIURL = "https://operations.cropwise.com/api/v3/fuel_types"

type FuelTypesResponse struct {
	Data []models.FuelTypeModel `json:"data"`
}

// FetchAndSaveFuelTypes - функция для загрузки данных типов топлива и сохранения их в базу данных
func FetchAndSaveFuelTypes(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FuelTypesAPIURL, nil)
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
		return errors.New("не удалось получить данные о типах топлива")
	}

	// Декодируем ответ
	var fuelTypesResponse FuelTypesResponse
	if err := json.NewDecoder(resp.Body).Decode(&fuelTypesResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, fuelType := range fuelTypesResponse.Data {
		if err := database.DB.Create(&fuelType).Error; err != nil {
			return fmt.Errorf("ошибка сохранения типа топлива с ID %d: %w", fuelType.ID, err)
		}
	}

	return nil
}
