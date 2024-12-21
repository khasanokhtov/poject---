package models

import (
	"time"
)

type Fertilizer struct {
	ID                    int64      `gorm:"primaryKey;column:id" json:"id"`
	Name                  *string    `gorm:"column:name;size:255" json:"name"`
	FertilizerType        *string    `gorm:"column:fertilizer_type;size:255" json:"fertilizer_type"`
	SourceType            *string    `gorm:"column:source_type;size:255" json:"source_type"`
	NutrientType          *string    `gorm:"column:nutrient_type;size:255" json:"nutrient_type"`
	UnitsOfMeasurement    *string    `gorm:"column:units_of_measurement;size:255" json:"units_of_measurement"`
	Archived              *bool      `gorm:"column:archived;default:false" json:"archived"`
	WhItemID              *int64     `gorm:"column:wh_item_id" json:"wh_item_id"`
	WhItemBaseUnitID      *int64     `gorm:"column:wh_item_base_unit_id" json:"wh_item_base_unit_id"`
	AdditionalInfo        *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	Description           *string    `gorm:"column:description;type:text" json:"description"`
	Density               *string    `gorm:"column:density;size:255" json:"density"`
	ExternalID            *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt             *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt             *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Fertilizer) TableName() string {
	return "fertilizers"
}