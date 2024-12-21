package models

import (
	"time"
)

type Seed struct {
	ID                     uint       `gorm:"primaryKey;column:id" json:"id"`
	Name                   *string    `gorm:"column:name;size:255" json:"name"`
	CropID                 *uint      `gorm:"column:crop_id" json:"crop_id"`
	AdditionalInfo         *string    `gorm:"column:additional_info;size:255" json:"additional_info"`
	Archived               *bool      `gorm:"column:archived" json:"archived"`
	Description            *string    `gorm:"column:description;size:255" json:"description"` 
	UnitsOfMeasurement     *string    `gorm:"column:units_of_measurement;size:255" json:"units_of_measurement"`
	Variety                *string    `gorm:"column:variety;size:255" json:"variety"`
	ReproductionNumber     *int       `gorm:"column:reproduction_number" json:"reproduction_number"`
	Reproduction           *string    `gorm:"column:reproduction;size:255" json:"reproduction"`
	WhItemID               *uint      `gorm:"column:wh_item_id" json:"wh_item_id"`
	WhItemBaseUnitID       *uint      `gorm:"column:wh_item_base_unit_id" json:"wh_item_base_unit_id"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`
	ExternalID             *string    `gorm:"column:external_id;size:255" json:"external_id"`
	RipenessGroup          *string    `gorm:"column:ripeness_group;size:255" json:"ripeness_group"`
	RipenessGroupName      *string    `gorm:"column:ripeness_group_name;size:255" json:"ripeness_group_name"`
	ThousandSeedsMass      *float64   `gorm:"column:thousand_seeds_mass" json:"thousand_seeds_mass"`
	SowingSuitability      *string    `gorm:"column:sowing_suitability;size:255" json:"sowing_suitability"`
	PiecesPerOneSowingUnit *float64   `gorm:"column:pieces_per_one_sowing_unit" json:"pieces_per_one_sowing_unit"`
}

func (Seed) TableName() string {
	return "seeds"
}
