package models

import (
	"time"
)

type FuelTypeModel struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	ShortName    *string   `json:"short_name"`
	StandardName string    `json:"standard_name"`
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ExternalID   *string   `json:"external_id"`
}
