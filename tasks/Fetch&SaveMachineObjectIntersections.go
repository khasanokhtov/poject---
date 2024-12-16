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

const MachineObjectIntersectionsAPIURL = "https://operations.cropwise.com/api/v3/machine_object_intersections"

type MachineObjectIntersectionsResponse struct {
	Data []models.MachineObjectIntersectionModel `json:"data"`
	Meta struct {
		Response struct {
			ObtainedRecords int `json:"obtained_records"`
			LastRecordID    int `json:"last_record_id"`
		} `json:"response"`
	} `json:"meta"`
}

// FetchAndSaveMachineObjectIntersections - функция для загрузки и сохранения данных пересечений объектов машин
func FetchAndSaveMachineObjectIntersections(token string, schemaName string) error {
	log.Printf("Начинаем загрузку данных MachineObjectIntersections для схемы: %s", schemaName)

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
		url := MachineObjectIntersectionsAPIURL + "?from_id=" + strconv.Itoa(fromID)
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

		var response MachineObjectIntersectionsResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}

		// Сохраняем данные в базу данных
		for _, intersection := range response.Data {
			if err := database.DB.Save(&intersection).Error; err != nil {
				return fmt.Errorf("ошибка сохранения пересечения с ID %d: %w", intersection.ID, err)
			}
		}

		// Проверяем, нужно ли продолжать загрузку данных
		if response.Meta.Response.ObtainedRecords == 0 {
			break
		}

		fromID = response.Meta.Response.LastRecordID + 1
	}

	return nil
}
