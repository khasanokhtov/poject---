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

const AgronomistAssignmentsEndpoint = "https://operations.cropwise.com/api/v3/agronomist_assignments"


func FetchAndSaveAgronomistAssignments(token, schemaName string) error {
	log.Printf("Начинаем загрузку данных для схемы: %s", schemaName)

	// Устанавливаем search_path на схему компании
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

	// Логируем текущий search_path
	var currentSearchPath string
	if err := database.DB.Raw("SHOW search_path").Scan(&currentSearchPath).Error; err != nil {
		log.Printf("Ошибка получения search_path: %v", err)
	} else {
		log.Printf("Текущий search_path: %s", currentSearchPath)
	}

	for {
		url := AgronomistAssignmentsEndpoint + "?from_id=" + strconv.Itoa(fromID)
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

		var response struct {
			Data []models.AgronomistAssignmentModel `json:"data"`
			Meta struct {
				Response struct {
					ObtainedRecords int `json:"obtained_records"`
					LastRecordID    int `json:"last_record_id"`
				} `json:"response"`
			} `json:"meta"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %w", err)
		}

		for _, assignment := range response.Data {
			if err := database.DB.Save(&assignment).Error; err != nil {
				return fmt.Errorf("ошибка сохранения агрономического назначения ID %d: %w", assignment.ID, err)
			}
		}

		if response.Meta.Response.ObtainedRecords == 0 {
			break
		}

		fromID = response.Meta.Response.LastRecordID + 1
	}

	log.Printf("Данные успешно загружены для схемы %s", schemaName)
	return nil
}
