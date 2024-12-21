package repository

import (
	"gorm.io/gorm"
)

// PlanFactDomain - структура для хранения данных отчета
type PlanFactDomain struct {
	Region        string  `json:"region"`
	Crop          string  `json:"crop"`
	WorkType      string  `json:"work_type"`
	Subtype       string  `json:"subtype"`
	PlannedArea   float64 `json:"planned_area"`
	FactArea      float64 `json:"fact_area"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
}

// PlanFactRepository - интерфейс для работы с отчетом
type PlanFactRepository interface {
	GetPlanFactTable(schema string) ([]PlanFactDomain, error)
}

// PlanFactRepositoryBase - реализация репозитория
type PlanFactRepositoryBase struct {
	DB *gorm.DB
}

// NewPlanFactRepository - конструктор
func NewPlanFactRepository(db *gorm.DB) PlanFactRepository {
	return &PlanFactRepositoryBase{DB: db}
}

// GetPlanFactTable - метод для извлечения данных отчета план-факт
func (r *PlanFactRepositoryBase) GetPlanFactTable(schema string) ([]PlanFactDomain, error) {
	var results []PlanFactDomain

	tx := r.DB.Table(schema + ".agro_operations ao")
	tx.Select(`
		fg.name as region,
		w.name as work_type,
		w.description as subtype,
		COALESCE(ao.planned_area, 0) as planned_area,
		COALESCE(ao.completed_area, 0) as fact_area,
		ao.planned_start_date as start_date,
		ao.planned_end_date as end_date
	`).
		Joins("JOIN " + schema + ".fields f ON ao.field_id = f.id").
		Joins("JOIN " + schema + ".field_groups fg ON f.field_group_id = fg.id").
		Joins("JOIN " + schema + ".work_types w ON ao.work_type_id = w.id")

	if err := tx.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
