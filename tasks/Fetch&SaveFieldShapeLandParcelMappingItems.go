package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const FieldShapeLandParcelMappingItemsAPIURL = "https://operations.cropwise.com/api/v3/field_shape_land_parcel_mapping_items"

type FieldShapeLandParcelMappingItemsResponse struct {
	Data []models.FieldShapeLandParcelMappingItemModel `json:"data"`
}

func FetchAndSaveFieldShapeLandParcelMappingItems(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", FieldShapeLandParcelMappingItemsAPIURL, nil)
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
		return errors.New("не удалось получить данные о сопоставлении участков земель и форм полей")
	}

	// Декодируем ответ
	var mappingItemsResponse FieldShapeLandParcelMappingItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&mappingItemsResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, mappingItem := range mappingItemsResponse.Data {
		if err := database.DB.Create(&mappingItem).Error; err != nil {
			return fmt.Errorf("ошибка сохранения записи с ID %d: %w", mappingItem.ID, err)
		}
	}

	return nil
}
