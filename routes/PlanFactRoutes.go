package routes

import (
	"integration-cropwise-v1/controller"
	"integration-cropwise-v1/middlewares"
	"integration-cropwise-v1/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupPlanFactRoutes - маршруты для отчета план-факт
func SetupPlanFactRoutes(app *fiber.App, db *gorm.DB) {
	repo := repository.NewPlanFactRepository(db)
	ctrl := controller.NewPlanFactController(repo)

	app.Get("/get-plan-fact", middlewares.AuthMiddleware, ctrl.GetPlanFact)
}
