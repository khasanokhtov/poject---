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

    // Отправляем запрос на внешнюю API
    authResponse, err := services.AuthenticateUser(authRequest.Email, authRequest.Password, 0) // localCompanyID передается позже
    if err != nil {
        log.Printf("Ошибка аутентификации через API: %v", err)
        return c.Status(http.StatusUnauthorized).JSON(AuthResponse{
            Message: "Authentication failed",
        })
    }

    // Проверяем, есть ли компания в локальной базе по company_id
    var company models.Company
    if err := database.DB.Where("company_id = ?", authResponse.CompanyID).First(&company).Error; err != nil {
        log.Printf("Компания с company_id %d не найдена в локальной базе: %v", authResponse.CompanyID, err)
        return c.Status(http.StatusForbidden).JSON(AuthResponse{
            Message: "Company not authorized",
        })
    }

    // Если всё успешно, возвращаем успешный ответ
    return c.Status(http.StatusOK).JSON(AuthResponse{
        Token:   authResponse.UserApiToken,
        Message: "Authentication successful",
    })
}
