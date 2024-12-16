package models

import "time"

// MaintenanceType - модель для таблицы maintenance_types
type MaintenanceType struct {
	ID                     int        `json:"id" gorm:"primaryKey"`
	MaintenanceTypeGroupID int        `json:"maintenance_type_group_id"`
	Name                   string     `json:"name"`
	ExternalID             *string    `json:"external_id"`  // Может быть null
	Description            *string    `json:"description"`  // Может быть null
	Hidden                 bool       `json:"hidden"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}
