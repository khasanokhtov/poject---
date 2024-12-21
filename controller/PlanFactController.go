package controller

import (
	"integration-cropwise-v1/repository"

	"github.com/gofiber/fiber/v2"
)

// PlanFactController - структура контроллера
type PlanFactController struct {
	Repo repository.PlanFactRepository
}

// NewPlanFactController - конструктор контроллера
func NewPlanFactController(repo repository.PlanFactRepository) *PlanFactController {
	return &PlanFactController{Repo: repo}
}

// GetPlanFact - обработчик для получения отчета план-факт
func (c *PlanFactController) GetPlanFact(ctx *fiber.Ctx) error {
	schema := ctx.Query("schema")
	if schema == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "schema is required"})
	}

	results, err := c.Repo.GetPlanFactTable(schema)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	plannedArea := 0.0
	factArea := 0.0
	for _, res := range results {
		plannedArea += res.PlannedArea
		factArea += res.FactArea
	}

	overallProgress := 0.0
	if plannedArea > 0 {
		overallProgress = (factArea / plannedArea) * 100
	}

	headerCards := []fiber.Map{
		{"title": "План", "value": plannedArea, "color": "#00FF00"},
		{"title": "Факт", "value": factArea, "color": "#800080"},
		{"title": "Прогресс", "value": overallProgress, "color": "#808080"},
	}

	return ctx.JSON(fiber.Map{
		"headerCards": headerCards,
		"planFactTable": results,
	})
}
