package models

import "time"

// MaintenanceRecord - модель для таблицы maintenance_records
type MaintenanceRecord struct {
	ID                int        `json:"id" gorm:"primaryKey"`
	MaintainableID    int        `json:"maintainable_id"`
	MaintainableType  string     `json:"maintainable_type"`
	MaintenancePlanID *int       `json:"maintenance_plan_id"` // Может быть null
	Status            string     `json:"status"`
	Season            int        `json:"season"`
	Motohours         *float64   `json:"motohours"`           // Может быть null
	Mileage           *float64   `json:"mileage"`             // Может быть null
	StartTime         time.Time  `json:"start_time"`
	EndTime           time.Time  `json:"end_time"`
	PlannedStartTime  time.Time  `json:"planned_start_time"`
	PlannedEndTime    time.Time  `json:"planned_end_time"`
	Description       *string    `json:"description"`         // Может быть null
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
