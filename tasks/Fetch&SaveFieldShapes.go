package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FieldShapesAPIURL = "https://operations.cropwise.com/api/v3/field_shapes"

type FieldShapesResponse struct {
	Data []models.FieldShapeModel `json:"data"`
}

func FetchAndSaveFieldShapes(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FieldShapesAPIURL, nil)
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
		return errors.New("не удалось получить данные о формах полей")
	}

	// Декодируем ответ
	var fieldShapesResponse FieldShapesResponse
	if err := json.NewDecoder(resp.Body).Decode(&fieldShapesResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, fieldShape := range fieldShapesResponse.Data {
		if err := database.DB.Create(&fieldShape).Error; err != nil {
			return fmt.Errorf("ошибка сохранения записи с ID %d: %w", fieldShape.ID, err)
		}
	}

	return nil
}
