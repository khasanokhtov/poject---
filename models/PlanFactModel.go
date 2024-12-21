package models

import "time"

// PlanFact - модель для хранения данных отчета план-факт
type PlanFact struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Region        string    `json:"region"`
	Crop          string    `json:"crop"`
	WorkType      string    `json:"work_type"`
	Subtype       string    `json:"subtype"`
	PlannedArea   float64   `json:"planned_area"`
	FactArea      float64   `json:"fact_area"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}