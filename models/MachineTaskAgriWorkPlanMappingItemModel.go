package models

import (
	"time"
)

// MachineTaskAgriWorkPlanMappingItemModel - структура для хранения данных связи задач машин и агропланов
type MachineTaskAgriWorkPlanMappingItemModel struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	MachineTaskID  uint      `json:"machine_task_id"`
	AgriWorkPlanID uint      `json:"agri_work_plan_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
