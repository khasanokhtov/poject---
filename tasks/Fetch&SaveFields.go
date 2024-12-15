package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FieldsAPIURL = "https://operations.cropwise.com/api/v3a/fields"

type FieldsResponse struct {
	Data []models.FieldModel `json:"data"`
}

// FetchAndSaveFields - функция для загрузки данных полей и сохранения их в базу данных
func FetchAndSaveFields(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FieldsAPIURL, nil)
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
		return errors.New("не удалось получить данные о полях")
	}

	// Декодируем ответ
	var fieldsResponse FieldsResponse
	if err := json.NewDecoder(resp.Body).Decode(&fieldsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, field := range fieldsResponse.Data {
		if err := database.DB.Create(&field).Error; err != nil {
			return fmt.Errorf("ошибка сохранения поля с ID %d: %w", field.ID, err)
		}
	}

	return nil
}
