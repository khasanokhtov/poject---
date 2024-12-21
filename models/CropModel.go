package models

import (
	"time"
)

type Crop struct {
	ID                         int64      `gorm:"primaryKey;column:id" json:"id"`
	Name                       string     `gorm:"column:name;size:255" json:"name"`
	ShortName                  *string    `gorm:"column:short_name;size:255" json:"short_name"`
	StandardName               *string    `gorm:"column:standard_name;size:255" json:"standard_name"`
	Color                      *string    `gorm:"column:color;size:255" json:"color"`
	AdditionalInfo             *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	Description                *string    `gorm:"column:description;type:text" json:"description"`
	SeasonType                 *string    `gorm:"column:season_type;size:255" json:"season_type"`
	BaseCropID                 *int64     `gorm:"column:base_crop_id"  json:"base_crop_id"`
	ProductivityEstimateCropName *string  `gorm:"column:productivity_estimate_crop_name;size:255" json:"productivity_estimate_crop_name"`
	PlantThreatCropName        *string    `gorm:"column:plant_threat_crop_name;size:255"  json:"plant_threat_crop_name"`
	Hidden                     *bool      `gorm:"column:hidden;default:false" json:"hidden"`
	ExternalID                 *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt                  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Crop) TableName() string {
	return "crops"
}