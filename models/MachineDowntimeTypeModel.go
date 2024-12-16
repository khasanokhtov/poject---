package models

import (
	"time"
)

// MachineDowntimeTypeModel - структура для хранения данных о типах простоев машин
type MachineDowntimeTypeModel struct {
	ID                          uint       `json:"id" gorm:"primaryKey"`
	Name                        string     `json:"name"`
	CustomName                  *string    `json:"custom_name"`
	StandardName                string     `json:"standard_name"`
	MachineDowntimeTypeGroupID  uint       `json:"machine_downtime_type_group_id"`
	AdditionalInfo              *string    `json:"additional_info"`
	Hidden                      *bool      `json:"hidden"`
	CreatedAt                   time.Time  `json:"created_at"`
	UpdatedAt                   time.Time  `json:"updated_at"`
	ExternalID                  *string    `json:"external_id"`
}
