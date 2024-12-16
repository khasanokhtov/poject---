package models

import "time"

// MachineWorkPlanRowModel - модель для таблицы machine_work_plan_rows
type MachineWorkPlanRowModel struct {
	ID                int        `json:"id" gorm:"primaryKey"`
	MachineWorkPlanID int        `json:"machine_work_plan_id"`
	MachineID         int        `json:"machine_id"`
	ImplementID       *int       `json:"implement_id"` // Может быть null
	Ind               int        `json:"ind"`
	Rate              float64    `json:"rate"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
