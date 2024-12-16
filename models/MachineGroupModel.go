package models

import (
	"time"
)

// MachineGroupModel - структура для хранения данных о группах машин
type MachineGroupModel struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null"`
	AdditionalInfo string    `json:"additional_info"`
	Description    string    `json:"description"`
	ExternalID     *string   `json:"external_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
