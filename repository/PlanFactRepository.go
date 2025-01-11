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

type PlanFactFilters struct{
	StartDate 	 string
	EndDate   	 string
	Crop         string
	WorkType   	 string
	WorkSubtype  string
	Region    	 string
	Limit		 int
	Offset		 int
}

// PlanFactRepository - интерфейс для работы с отчетом
type PlanFactRepository interface {
	GetPlanFactTable(schema string, filters PlanFactFilters) ([]PlanFactDomain, error)
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
func (r *PlanFactRepositoryBase) GetPlanFactTable(schema string, filters PlanFactFilters) ([]PlanFactDomain, error) {
	var results []PlanFactDomain

	tx := r.DB.Table(schema + ".agro_operations ao")
	tx.Select(`
	fg.name as region,
	cr.name as crop,
	w.name as work_type,
	w.description as subtype,
	COALESCE(ao.planned_area, 0) as planned_area,
	COALESCE(ao.completed_area, 0) as fact_area,
	ao.planned_start_date as start_date,
	ao.planned_end_date as end_date
`).
	Joins("JOIN " + schema + ".fields f ON ao.field_id = f.id").
	Joins("JOIN " + schema + ".field_groups fg ON f.field_group_id = fg.id").
	Joins("JOIN " + schema + ".work_types w ON ao.work_type_id = w.id").
	Joins("LEFT JOIN " + schema + ".history_items hi ON ao.field_id = hi.field_id").
	Joins("LEFT JOIN " + schema + ".crops cr ON hi.crop_id = cr.id")

// Применение фильтров
if filters.StartDate != "" && filters.EndDate != "" {
	tx = tx.Where("ao.planned_start_date BETWEEN ? AND ?", filters.StartDate, filters.EndDate)
}
if filters.WorkType != "" {
	tx = tx.Where("w.name = ?", filters.WorkType)
}
if filters.WorkSubtype != "" {
	tx = tx.Where("w.description = ?", filters.WorkSubtype)
}
if filters.Crop != "" {
	tx = tx.Where("cr.name = ?", filters.Crop)
}
if filters.Region != "" {
	tx = tx.Where("fg.name = ?", filters.Region)
}

// Пагинация
if filters.Limit > 0 {
    tx = tx.Limit(filters.Limit)
} else {
    tx = tx.Limit(50) // Значение по умолчанию
}
if filters.Offset > 0 {
	tx = tx.Offset(filters.Offset)
}

// Выполнение запроса
if err := tx.Find(&results).Error; err != nil {
	return nil, err
}

	return results, nil
}
