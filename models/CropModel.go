package models

import (
	"time"
)

type Crop struct {
	ID                         int64      `gorm:"primaryKey;column:id"`
	Name                       string     `gorm:"column:name;size:255"`
	ShortName                  *string    `gorm:"column:short_name;size:255"`
	StandardName               *string    `gorm:"column:standard_name;size:255"`
	Color                      *string    `gorm:"column:color;size:255"`
	AdditionalInfo             *string    `gorm:"column:additional_info;type:text"`
	Description                *string    `gorm:"column:description;type:text"`
	SeasonType                 *string    `gorm:"column:season_type;size:255"`
	BaseCropID                 *int64     `gorm:"column:base_crop_id"`
	ProductivityEstimateCropName *string   `gorm:"column:productivity_estimate_crop_name;size:255"`
	PlantThreatCropName        *string    `gorm:"column:plant_threat_crop_name;size:255"`
	Hidden                     *bool      `gorm:"column:hidden;default:false"`
	ExternalID                 *string    `gorm:"column:external_id;size:255"`
	CreatedAt                  *time.Time `gorm:"column:created_at"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at"`
}

func (Crop) TableName() string {
	return "crops"
}