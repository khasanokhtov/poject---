package routes

import (
	"integration-cropwise-v1/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/login", handlers.AuthenticateUser)
}