package handlers

import (
	"integration-cropwise-v1/models"
	"integration-cropwise-v1/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

//Обработчик для получения данных
func CreateCompany(c *fiber.Ctx) error {
	//Парсим тело
	var credentials models.CompanyCredentials
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	//Проверка валидности
	if credentials.Email == "" || credentials.Password == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email и пароль обязательны",
		})
	}

	log.Printf("Получены учетные данные: Email: %s, Password: %s", credentials.Email, credentials.Password)

	// Получаем токен
	authResponse, err := services.GetAuthToken(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Ошибка аутентификации: " + err.Error(),
		})
	}

	//Сохраняем данные компании
	err = services.SaveCompany(credentials.Email, authResponse.UserApiToken, authResponse.Company)
	if err != nil{
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Ошибка сохранения компании: " + err.Error(),
		})
	}

	//Успешный ответ

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Компания успешно создана и схема добавлена",
		"token":   authResponse.UserApiToken,
		"company": authResponse.Company,
	})
}