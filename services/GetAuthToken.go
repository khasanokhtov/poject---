package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"integration-cropwise-v1/models"
	"io"
	"net/http"
)

const AuthEndpoint = "https://operations.cropwise.com/api/v3/sign_in"

//Отправляем запрос на аутентификацию
func GetAuthToken(email, password string) (models.AuthResponse, error){
	//Данные для запроса
	authRequest := models.AuthRequest{}
	authRequest.UserLogin.Email = email
	authRequest.UserLogin.Password = password

	jsonData, err := json.Marshal(authRequest)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("ошибка сериализации данных: %w", err)
	}

	//Выполнение запроса
	resp, err := http.Post(AuthEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	//Статус Чек
	if resp.StatusCode != http.StatusOK{
		body, _ := io.ReadAll(resp.Body)
		return models.AuthResponse{}, fmt.Errorf("ошибка аутентификации: %s", string(body))
	}

	//Обрабатываем JSON-Ответ
	var authResponse models.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("ошибка парсинга ответа: %w", err)
	}

	//Проверка успешности
	if !authResponse.Success {
		return models.AuthResponse{}, fmt.Errorf("ошибка аутентификации")
	}

	return authResponse, nil
}