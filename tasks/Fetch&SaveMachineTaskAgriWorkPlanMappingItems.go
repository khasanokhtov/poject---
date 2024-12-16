package tasks

import (
	"encoding/json"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

const MachineTaskAgriWorkPlanMappingItemsAPIURL = "https://operations.cropwise.com/api/v3/machine_task_agri_work_plan_mapping_items"

type MachineTaskAgriWorkPlanMappingItemsResponse struct {
	Data []models.MachineTaskAgriWorkPlanMappingItemModel `json:"data"`
	Meta struct {
		Response struct {
			ObtainedRecords int `json:"obtained_records"`
			LastRecordID    int `json:"last_record_id"`
		} `json:"response"`
	} `json:"meta"`
}

// FetchAndSaveMachineTaskAgriWorkPlanMappingItems - функция для загрузки и сохранения данных связи задач машин и агропланов
func FetchAndSaveMachineTaskAgriWorkPlanMappingItems(token string, schemaName string) error {
	log.Printf("Начинаем загрузку MachineTaskAgriWorkPlanMappingItems для схемы: %s", schemaName)

	// Устанавливаем search_path для схемы компании
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка установки search_path на %s: %v", schemaName, err)
		return err
	}
	defer func() {
		resetSearchPath := "SET search_path TO public"
		if err := database.DB.Exec(resetSearchPath).Error; err != nil {
			log.Printf("Ошибка сброса search_path на public: %v", err)
		}
	}()

	client := &http.Client{}
	fromID := 0

	for {
		url := MachineTaskAgriWorkPlanMappingItemsAPIURL + "?from_id=" + strconv.Itoa(fromID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("ошибка создания запроса: %w", err)
		}
		req.Header.Set("X-User-Api-Token", token)

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("ошибка выполнения запроса: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("ошибка ответа: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %w", err)
		}

		var mappingItemsResponse MachineTaskAgriWorkPlanMappingItemsResponse
		if err := json.Unmarshal(body, &mappingItemsResponse); err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}

		// Сохраняем данные в базу
		for _, mappingItem := range mappingItemsResponse.Data {
			if err := database.DB.Save(&mappingItem).Error; err != nil {
				return fmt.Errorf("ошибка сохранения связи задачи машины и агроплана с ID %d: %w", mappingItem.ID, err)
			}
		}

		// Проверяем, нужно ли продолжать загрузку данных
		if mappingItemsResponse.Meta.Response.ObtainedRecords == 0 {
			break
		}

		fromID = mappingItemsResponse.Meta.Response.LastRecordID + 1
	}

	return nil
}
