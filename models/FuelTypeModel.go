package models

import (
	"time"
)

type FuelType struct {
	ID           int64      `gorm:"primaryKey;column:id" json:"id"`
	Name         *string    `gorm:"column:name;size:255" json:"name"`
	ShortName    *string    `gorm:"column:short_name;size:255" json:"short_name"`
	StandardName *string    `gorm:"column:standard_name;size:255" json:"standard_name"`
	Category     *string    `gorm:"column:category;size:255" json:"category"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	ExternalID   *string    `gorm:"column:external_id;size:255" json:"external_id"`
}

func (FuelType) TableName() string {
	return "fuel_types"
}