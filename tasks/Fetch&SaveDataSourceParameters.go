package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const DataSourceParametersAPIURL = "https://operations.cropwise.com/api/v3/data_source_parameters"

type DataSourceParametersResponse struct {
	Data []models.DataSourceParameterModel `json:"data"`
}

func FetchAndSaveDataSourceParameters(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", DataSourceParametersAPIURL, nil)
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
		return errors.New("не удалось получить данные о параметрах источника GPS")
	}

	// Декодируем ответ
	var dataSourceParametersResponse DataSourceParametersResponse
	if err := json.NewDecoder(resp.Body).Decode(&dataSourceParametersResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, parameter := range dataSourceParametersResponse.Data {
		if err := database.DB.Save(&parameter).Error; err != nil {
			return fmt.Errorf("ошибка сохранения параметра с ID %d: %w", parameter.ID, err)
		}
	}

	return nil
}
