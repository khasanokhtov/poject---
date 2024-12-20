package models

import (
	"time"
)

type Fertilizer struct {
	ID                    int64      `gorm:"primaryKey;column:id"`
	Name                  *string    `gorm:"column:name;size:255"`
	FertilizerType        *string    `gorm:"column:fertilizer_type;size:255"`
	SourceType            *string    `gorm:"column:source_type;size:255"`
	NutrientType          *string    `gorm:"column:nutrient_type;size:255"`
	UnitsOfMeasurement    *string    `gorm:"column:units_of_measurement;size:255"`
	Archived              *bool      `gorm:"column:archived;default:false"`
	WhItemID              *int64     `gorm:"column:wh_item_id"`
	WhItemBaseUnitID      *int64     `gorm:"column:wh_item_base_unit_id"`
	AdditionalInfo        *string    `gorm:"column:additional_info;type:text"`
	Description           *string    `gorm:"column:description;type:text"`
	Density               *string    `gorm:"column:density;size:255"`
	ExternalID            *string    `gorm:"column:external_id;size:255"`
	CreatedAt             *time.Time `gorm:"column:created_at"`
	UpdatedAt             *time.Time `gorm:"column:updated_at"`
}

func (Fertilizer) TableName() string {
	return "fertilizers"
}