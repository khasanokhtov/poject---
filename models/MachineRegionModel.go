package models

import (
	"time"
)

// MachineRegionModel - структура для хранения данных о регионах машин
type MachineRegionModel struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Ancestry      string    `json:"ancestry"`
	Description   string    `json:"description"`
	AdditionalInfo string   `json:"additional_info"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ExternalID    *string   `json:"external_id"`
}
