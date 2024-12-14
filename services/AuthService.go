package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	ExternalAuthURL         = "https://operations.cropwise.com/api/v3/sign_in"
	ExternalCompanyInfoURL  = "https://operations.cropwise.com/api/v3/company"
)

type ExternalAuthRequest struct {
	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user_login"`
}

type ExternalAuthResponse struct {
	Success      bool   `json:"success"`
	UserApiToken string `json:"user_api_token"`
	Company      string `json:"company"`
	CompanyID	 uint	`json:"-"`
}

type ExternalCompanyInfoToCheck struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ExternalCompanyAuthResponse struct {
	Data ExternalCompanyInfoToCheck `json:"data"`
}

func AuthenticateUser(email, password string, localCompanyID uint) (*ExternalAuthResponse, error) {
	// 1. Аутентификация пользователя и получение токена
	authResponse, err := authenticate(email, password)
	if err != nil {
		return nil, fmt.Errorf("ошибка аутентификации: %w", err)
	}

	// 2. Получение информации о компании
	companyID, err := fetchCompanyID(authResponse.UserApiToken)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения ID компании: %w", err)
	}

	
	authResponse.CompanyID = companyID
	return authResponse, nil
}

func authenticate(email, password string) (*ExternalAuthResponse, error) {
	// Формируем запрос
	authRequest := ExternalAuthRequest{}
	authRequest.UserLogin.Email = email
	authRequest.UserLogin.Password = password

	jsonData, err := json.Marshal(authRequest)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	// Отправляем запрос
	resp, err := http.Post(ExternalAuthURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("внешняя аутентификация не удалась")
	}

	// Обрабатываем ответ
	var authResponse ExternalAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	if !authResponse.Success {
		return nil, errors.New("аутентификация не удалась")
	}

	return &authResponse, nil
}

func fetchCompanyID(userApiToken string) (uint, error) {
	req, err := http.NewRequest("GET", ExternalCompanyInfoURL, nil)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	req.Header.Set("X-User-Api-Token", userApiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("не удалось получить информацию о компании")
	}

	var companyResponse ExternalCompanyAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&companyResponse); err != nil {
		return 0, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return companyResponse.Data.ID, nil
}