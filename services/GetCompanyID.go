package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)


const GetCompanyUrlInfo = "https://operations.cropwise.com/api/v3/company"

type ExternalCompanyInfo struct{
	ID uint `json:"id"`
	Name string `json:"name"`
}


type ExternalCompanyResponse struct {
	Data ExternalCompanyInfo `json:"data"`
}

func FetchCompanyID(userApiToken string) (uint, error) {
	req, err := http.NewRequest("GET", GetCompanyUrlInfo, nil)
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

	if resp.StatusCode != http.StatusOK{
		return 0, errors.New("не удалось получить информацию о компании")
	}

	var companyResponse ExternalCompanyResponse
	if err := json.NewDecoder(resp.Body).Decode(&companyResponse); err != nil {
		return 0, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return companyResponse.Data.ID, nil
}