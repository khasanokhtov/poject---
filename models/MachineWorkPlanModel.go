package models

import "time"

// MachineWorkPlanModel - модель для таблицы machine_work_plans
type MachineWorkPlanModel struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	AgriWorkPlanID  int       `json:"agri_work_plan_id"`
	Color           string    `json:"color"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
