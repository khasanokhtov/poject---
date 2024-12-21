package models

import (
	"time"
)

type HistoryItem struct {
	ID                    int64      `gorm:"primaryKey;column:id" json:"id"`
	FieldID               *int64     `gorm:"column:field_id" json:"field_id"`
	Year                  *int64     `gorm:"column:year" json:"year"`
	Active                *bool      `gorm:"column:active" json:"active"`
	CropID                *int64     `gorm:"column:crop_id" json:"crop_id"`
	Variety               *string    `gorm:"column:variety;size:255" json:"variety"`
	TillType              *string    `gorm:"column:till_type;size:255" json:"till_type"`
	Productivity          *float64   `gorm:"column:productivity" json:"productivity"`
	ProductivityEstimate  *float64   `gorm:"column:productivity_estimate" json:"productivity_estimate"`
	ProductivityZone      *string    `gorm:"column:productivity_zone;size:255" json:"productivity_zone"`
	SowingDate            *string    `gorm:"column:sowing_date" json:"sowing_date"`
	HarvestingDate        *string    `gorm:"column:harvesting_date" json:"harvesting_date"`
	Description           *string    `gorm:"column:description;type:text" json:"description"`
	AdditionalInfo        *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	ExternalID            *string    `gorm:"column:external_id;size:255" json:"external_id"`
	HarvestedWeight       *float64   `gorm:"column:harvested_weight" json:"harvested_weight"`
	MarketableWeight      *float64   `gorm:"column:marketable_weight" json:"marketable_weight"`
	YieldDensity          *float64   `gorm:"column:yield_density" json:"yield_density"`
	ExpectedYield         *float64   `gorm:"column:expected_yield" json:"expected_yield"`
	AgeOfSugarCane        *float64   `gorm:"column:age_of_sugar_cane" json:"age_of_sugar_cane"`
	GrainClass            *string    `gorm:"column:grain_class;size:255" json:"grain_class"`
	GrainHumidity         *float64   `gorm:"column:grain_humidity" json:"grain_humidity"`
	IrrigationType        *string    `gorm:"column:irrigation_type;size:255" json:"irrigation_type"`
	GrainGarbageAdmixture *float64   `gorm:"column:grain_garbage_admixture" json:"grain_garbage_admixture"`
	CreatedAt             *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt             *time.Time `gorm:"column:updated_at" json:"updated_at"`
	ProductionCycleID     *int64     `gorm:"column:production_cycle_id" json:"production_cycle_id"`
	AutomaticEndDate      *time.Time `gorm:"column:automatic_end_date" json:"automatic_end_date"`
	AutoShapeDetect       *bool      `gorm:"column:auto_shape_detect" json:"auto_shape_detect"`
	FieldShapeID          *int64     `gorm:"column:field_shape_id" json:"field_shape_id"`
}

func (HistoryItem) TableName() string {
	return "history_items"
}