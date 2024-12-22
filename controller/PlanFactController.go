package controller

import (
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"integration-cropwise-v1/repository"
	"log"
	"math"

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
	// Извлекаем companyID из middleware
	companyID, ok := ctx.Locals("companyID").(uint)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Проверяем наличие компании в базе
	var company models.Company
	if err := database.DB.Where("company_id = ?", companyID).First(&company).Error; err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Компания не найдена в базе"})
	}

	// Определяем схему из имени компании
	schema, err := findSchemaByCompanyID(companyID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Получаем данные отчета
	results, err := c.Repo.GetPlanFactTable(schema)
	if err != nil {
		log.Printf("Ошибка при получении данных отчета: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Расчеты
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
	overallProgress = math.Round(overallProgress*100) / 100

	// Подготовка ответа
	headerCards := []fiber.Map{
		{"title": "План", "value": plannedArea, "color": "#00FF00"},
		{"title": "Факт", "value": factArea, "color": "#800080"},
		{"title": "Прогресс", "value": overallProgress, "color": "#808080"},
	}

	return ctx.JSON(fiber.Map{
		"headerCards":   headerCards,
		"planFactTable": results,
	})
}

func findSchemaByCompanyID(companyID uint) (string, error) {
    var schemaName string

    query := `
        SELECT schema_name
        FROM companies
        WHERE company_id = $1
    `

    err := database.DB.Raw(query, companyID).Scan(&schemaName).Error
    if err != nil || schemaName == "" {
        return "", fmt.Errorf("схема для компании с ID %d не найдена", companyID)
    }

    return schemaName, nil
}