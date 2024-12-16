package models

import (
	"time"
)

// MachineDowntimeTypeGroupModel - структура для хранения данных о группах типов простоев машин
type MachineDowntimeTypeGroupModel struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Name          string     `json:"name"`
	CustomName    *string    `json:"custom_name"`
	StandardName  string     `json:"standard_name"`
	AdditionalInfo *string   `json:"additional_info"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	ExternalID    *string    `json:"external_id"`
}
