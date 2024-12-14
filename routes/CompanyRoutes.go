package routes

import (
	"integration-cropwise-v1/handlers"

	"github.com/gofiber/fiber/v2"
)

//Маршрут для работы с компанией
func SetupCompanyRoutes(app *fiber.App){
	app.Post("/create-company", handlers.CreateCompany)
}