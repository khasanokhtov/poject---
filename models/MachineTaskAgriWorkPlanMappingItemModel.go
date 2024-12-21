package models

import (
	"time"
)

// MachineTaskAgriWorkPlanMappingItemModel - структура для хранения данных связи задач машин и агропланов
type MachineTaskAgriWorkPlanMappingItemModel struct {
	ID             uint      `gorm:"primaryKey;column:id" json:"id"`
	MachineTaskID  uint      `gorm:"column:machine_task_id" json:"machine_task_id"`
	AgriWorkPlanID uint      `gorm:"column:agri_work_plan_id" json:"agri_work_plan_id"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (MachineTaskAgriWorkPlanMappingItemModel) TableName() string {
	return "machines_task_agro_work_plan_mapping_item"
}
