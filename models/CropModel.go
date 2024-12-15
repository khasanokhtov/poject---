package models

import (
	"time"
)

type CropModel struct {
	ID                         uint      `json:"id" gorm:"primaryKey"`
	Name                       string    `json:"name" gorm:"not null"`
	ShortName                  *string   `json:"short_name"`
	StandardName               string    `json:"standard_name" gorm:"not null"`
	Color                      string    `json:"color"`
	AdditionalInfo             string    `json:"additional_info"`
	Description                string    `json:"description"`
	SeasonType                 string    `json:"season_type"`
	BaseCropID                 uint      `json:"base_crop_id"`
	ProductivityEstimateCropName string `json:"productivity_estimate_crop_name"`
	PlantThreatCropName        string    `json:"plant_threat_crop_name"`
	Hidden                     bool      `json:"hidden"`
	ExternalID                 *string   `json:"external_id"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}