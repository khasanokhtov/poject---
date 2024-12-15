package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"net/http"
)

const GroupFoldersAPIURL = "https://operations.cropwise.com/api/v3/group_folders"

type GroupFoldersResponse struct {
	Data []models.GroupFolderModel `json:"data"`
}

// FetchAndSaveGroupFolders - загрузка данных и сохранение в базу данных
func FetchAndSaveGroupFolders(token string, schemaName string) error {
	// Устанавливаем search_path для работы с нужной схемой
	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		return fmt.Errorf("ошибка переключения схемы: %w", err)
	}

	// Создаем запрос к внешнему API
	req, err := http.NewRequest("GET", GroupFoldersAPIURL, nil)
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
		return errors.New("не удалось получить данные о Group Folders")
	}

	// Декодируем ответ
	var groupFoldersResponse GroupFoldersResponse
	if err := json.NewDecoder(resp.Body).Decode(&groupFoldersResponse); err != nil {
		return fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	// Сохраняем данные в базу
	for _, folder := range groupFoldersResponse.Data {
		if err := database.DB.Create(&folder).Error; err != nil {
			return fmt.Errorf("ошибка сохранения Group Folder с ID %d: %w", folder.ID, err)
		}
	}

	return nil
}
