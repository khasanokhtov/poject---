package handlers

import (
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"integration-cropwise-v1/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message"`
}

func AuthenticateUser(c *fiber.Ctx) error {
	var authRequest AuthRequest

	if err := c.BodyParser(&authRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(AuthResponse{
			Message: "Invalid request payload",
		})
	}

	var companyID models.Company
	if err := database.DB.Where("email = ?", authRequest.Email).First(&companyID).Error; err != nil {
		log.Printf("Компания с email %s не найдена в локальной базе: %v", authRequest.Email, err)
		return c.Status(http.StatusForbidden).JSON(AuthResponse{
			Message: "Company not authorized",
		})
	}
	localCompanyID := companyID.ID

	//Отправляем запрос на внешнюю API
	authResponse, err := services.AuthenticateUser(authRequest.Email, authRequest.Password, localCompanyID)
	if err != nil {
		log.Printf("Ошибка аутентификации через API: %v", err)
		return c.Status(http.StatusUnauthorized).JSON(AuthResponse{
			Message: "Atuhentication failed",
		})
	}

	// Проверяем, есть ли компания в локальной базе
	var company models.Company
	if err := database.DB.Where("schema_name = ?", authResponse.Company).First(&company).Error; err != nil {
		log.Printf("Компания с schema_name %s не найдена в локальной базе: %v", authResponse.Company, err)
		return c.Status(http.StatusForbidden).JSON(AuthResponse{
			Message: "Company not authorized",
		})
	}

	// Если всё успешно, возвращаем успешный ответ
	return c.Status(http.StatusOK).JSON(AuthResponse{
		Token: authResponse.UserApiToken,
		Message: "Authentication successful",
	})
}