package models

import "time"

// MaintenanceTypeGroup - модель для таблицы maintenance_type_groups
type MaintenanceTypeGroup struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	ExternalID  *string    `json:"external_id"`  // Может быть null
	Description *string    `json:"description"`  // Может быть null
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
