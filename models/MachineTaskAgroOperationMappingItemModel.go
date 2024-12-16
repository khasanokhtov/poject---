package models

import (
	"time"
)

// MachineTaskAgroOperationMappingItemModel - структура для хранения данных связи задач машин и агроопераций
type MachineTaskAgroOperationMappingItemModel struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	MachineTaskID  uint      `json:"machine_task_id"`
	AgroOperationID uint     `json:"agro_operation_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
