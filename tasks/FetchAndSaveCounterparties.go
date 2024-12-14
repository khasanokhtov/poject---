package tasks

import (
	"encoding/json"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"io"
	"log"
	"net/http"
)

const CounterpartiesEndpoint = "https://operations.cropwise.com/api/v3/counterparties"

func FetchAndSaveCounterparties(token, schemaName string) error {
	log.Printf("Начинаем загрузку данных для схемы: %s", schemaName)

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
	req, err := http.NewRequest("GET", CounterpartiesEndpoint, nil)
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
		Data []models.CounterpartyModel `json:"data"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	for _, counterparty := range response.Data {
		if err := database.DB.Save(&counterparty).Error; err != nil {
			return fmt.Errorf("ошибка сохранения контрагента ID %d: %w", counterparty.ID, err)
		}
	}

	log.Printf("Данные контрагентов успешно загружены для схемы %s", schemaName)
	return nil
}
