package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const EquipmentAssignmentsAPIURL = "https://operations.cropwise.com/api/v3/equipment_assignments"

type EquipmentAssignmentsResponse struct {
	Data []models.EquipmentAssignmentModel `json:"data"`
}

func FetchAndSaveEquipmentAssignments(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", EquipmentAssignmentsAPIURL, nil)
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
		return errors.New("не удалось получить данные о назначениях оборудования")
	}

	// Декодируем ответ
	var equipmentAssignmentsResponse EquipmentAssignmentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&equipmentAssignmentsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, assignment := range equipmentAssignmentsResponse.Data {
		if err := database.DB.Create(&assignment).Error; err != nil {
			return fmt.Errorf("ошибка сохранения назначения оборудования с ID %d: %w", assignment.ID, err)
		}
	}

	return nil
}
