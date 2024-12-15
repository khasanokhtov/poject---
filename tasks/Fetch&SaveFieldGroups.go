package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FieldGroupsAPIURL = "https://operations.cropwise.com/api/v3/field_groups"

type FieldGroupsResponse struct {
	Data []models.FieldGroupModel `json:"data"`
}

func FetchAndSaveFieldGroups(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FieldGroupsAPIURL, nil)
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
		return errors.New("не удалось получить данные о группах полей")
	}

	// Декодируем ответ
	var fieldGroupsResponse FieldGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&fieldGroupsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, fieldGroup := range fieldGroupsResponse.Data {
		if err := database.DB.Create(&fieldGroup).Error; err != nil {
			return fmt.Errorf("ошибка сохранения группы полей с ID %d: %w", fieldGroup.ID, err)
		}
	}

	return nil
}
