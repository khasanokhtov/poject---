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


const MachineEndpoint = "https://operations.cropwise.com/api/v3/machines"


func FetchAndSaveMachines(token, schemaName string) error {
	log.Printf("Начинаем загрузку данных для схемы: %s", schemaName)

	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка установки search_path на %s: %v", schemaName, err)
		return err
	}
	defer resetSearchPath()

	client := &http.Client{}
	fromID := 0

	for {
		url := MachineEndpoint + "?from_id=" + strconv.Itoa(fromID)
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
			Data []models.MachineModel `json:"data"`
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

		for _, machine := range response.Data {
			if err := database.DB.Save(&machine).Error; err != nil {
				return fmt.Errorf("ошибка сохранения машины ID %d: %w", machine.ID, err)
			}
		}

		if response.Meta.Response.ObtainedRecords == 0 {
			break
		}

		fromID = response.Meta.Response.LastRecordID + 1
	}

	return nil
}

func resetSearchPath() {
	resetSearchPath := "SET search_path TO public"
	if err := database.DB.Exec(resetSearchPath).Error; err != nil {
		log.Printf("Ошибка сброса search_path на public: %v", err)
	}
	log.Printf("search_path успешно сброшен на public")
}