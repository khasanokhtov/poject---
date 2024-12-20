package models

import (
	"time"
)

type FuelType struct {
	ID           int64      `gorm:"primaryKey;column:id"`
	Name         *string    `gorm:"column:name;size:255"`
	ShortName    *string    `gorm:"column:short_name;size:255"`
	StandardName *string    `gorm:"column:standard_name;size:255"`
	Category     *string    `gorm:"column:category;size:255"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
	ExternalID   *string    `gorm:"column:external_id;size:255"`
}

func (FuelType) TableName() string {
	return "fuel_types"
}