package models

import (
	"time"
)

// ImplementWorkTypeMappingItemModel - структура для привязки типов работ к навесному оборудованию
type ImplementWorkTypeMappingItemModel struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	ImplementID  uint      `json:"implement_id"`
	WorkTypeID   uint      `json:"work_type_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
