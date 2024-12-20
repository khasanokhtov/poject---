package models

import "time"

type MachineTaskAgroOperationMappingItem struct {
	ID              uint      `gorm:"primaryKey;column:id"`
	MachineTaskID   uint      `gorm:"column:machine_task_id"`
	AgroOperationID uint      `gorm:"column:agro_operation_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (MachineTaskAgroOperationMappingItem) TableName() string {
	return "machine_task_agro_operation_mapping_items"
}
