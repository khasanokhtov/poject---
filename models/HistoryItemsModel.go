package models

import (
	"time"
)

type HistoryItem struct {
	ID                    int64      `gorm:"primaryKey;column:id"`
	FieldID               *int64     `gorm:"column:field_id"`
	Year                  *int64     `gorm:"column:year"`
	Active                *bool      `gorm:"column:active"`
	CropID                *int64     `gorm:"column:crop_id"`
	Variety               *string    `gorm:"column:variety;size:255"`
	TillType              *string    `gorm:"column:till_type;size:255"`
	Productivity          *float64   `gorm:"column:productivity"`
	ProductivityEstimate  *float64   `gorm:"column:productivity_estimate"`
	ProductivityZone      *string    `gorm:"column:productivity_zone;size:255"`
	SowingDate            *time.Time `gorm:"column:sowing_date"`
	HarvestingDate        *time.Time `gorm:"column:harvesting_date"`
	Description           *string    `gorm:"column:description;type:text"`
	AdditionalInfo        *string    `gorm:"column:additional_info;type:text"`
	ExternalID            *string    `gorm:"column:external_id;size:255"`
	HarvestedWeight       *float64   `gorm:"column:harvested_weight"`
	MarketableWeight      *float64   `gorm:"column:marketable_weight"`
	YieldDensity          *float64   `gorm:"column:yield_density"`
	ExpectedYield         *float64   `gorm:"column:expected_yield"`
	AgeOfSugarCane        *float64   `gorm:"column:age_of_sugar_cane"`
	GrainClass            *string    `gorm:"column:grain_class;size:255"`
	GrainHumidity         *float64   `gorm:"column:grain_humidity"`
	IrrigationType        *string    `gorm:"column:irrigation_type;size:255"`
	GrainGarbageAdmixture *float64   `gorm:"column:grain_garbage_admixture"`
	CreatedAt             *time.Time `gorm:"column:created_at"`
	UpdatedAt             *time.Time `gorm:"column:updated_at"`
	ProductionCycleID     *int64     `gorm:"column:production_cycle_id"`
	AutomaticEndDate      *time.Time `gorm:"column:automatic_end_date"`
	AutoShapeDetect       *bool      `gorm:"column:auto_shape_detect"`
	FieldShapeID          *int64     `gorm:"column:field_shape_id"`
}

func (HistoryItem) TableName() string {
	return "history_items"
}