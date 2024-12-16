package models

import "time"

// WorkTypeGroup - модель для таблицы work_type_groups
type WorkTypeGroup struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	StandardName string     `json:"standard_name"`
	Description *string    `json:"description"` // Может быть null
	ExternalID  *string    `json:"external_id"` // Может быть null
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
