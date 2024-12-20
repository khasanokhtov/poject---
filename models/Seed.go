package models

import (
	"time"
)

type Seed struct {
	ID                     uint       `gorm:"primaryKey;column:id"`
	Name                   *string    `gorm:"column:name;size:255"`
	CropID                 *uint      `gorm:"column:crop_id"`
	AdditionalInfo         *string    `gorm:"column:additional_info;size:255"`
	Archived               *bool      `gorm:"column:archived"`
	Description            *string    `gorm:"column:description;size:255"`
	UnitsOfMeasurement     *string    `gorm:"column:units_of_measurement;size:255"`
	Variety                *string    `gorm:"column:variety;size:255"`
	ReproductionNumber     *int       `gorm:"column:reproduction_number"`
	Reproduction           *string    `gorm:"column:reproduction;size:255"`
	WhItemID               *uint      `gorm:"column:wh_item_id"`
	WhItemBaseUnitID       *uint      `gorm:"column:wh_item_base_unit_id"`
	CreatedAt              *time.Time `gorm:"column:created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at"`
	ExternalID             *string    `gorm:"column:external_id;size:255"`
	RipenessGroup          *string    `gorm:"column:ripeness_group;size:255"`
	RipenessGroupName      *string    `gorm:"column:ripeness_group_name;size:255"`
	ThousandSeedsMass      *float64   `gorm:"column:thousand_seeds_mass"`
	SowingSuitability      *string    `gorm:"column:sowing_suitability;size:255"`
	PiecesPerOneSowingUnit *float64   `gorm:"column:pieces_per_one_sowing_unit"`
}

func (Seed) TableName() string {
	return "seeds"
}
