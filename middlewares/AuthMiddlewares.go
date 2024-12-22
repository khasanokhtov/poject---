package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// External API для проверки токена
const ExternalCompanyInfoURL = "https://operations.cropwise.com/api/v3/company"

// ExtractCompanyID - функция для проверки токена
func ExtractCompanyID(token string) (uint, error) {
	req, err := http.NewRequest("GET", ExternalCompanyInfoURL, nil)
	if err != nil {
		return 0, errors.New("ошибка создания запроса")
	}

	req.Header.Set("X-User-Api-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New("ошибка выполнения запроса")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("недействительный токен")
	}

	// Обработка ответа
	var response struct {
		Data struct {
			ID uint `json:"id"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, errors.New("ошибка декодирования ответа")
	}

	return response.Data.ID, nil
}

// AuthMiddleware - добавляет companyID в контекст запроса
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Токен отсутствует"})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Неверный формат токена"})
	}

	companyID, err := ExtractCompanyID(tokenParts[1])
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Locals("companyID", companyID)
	return c.Next()
}
