package models

import (
	"time"
)

type MachineRegionMappingItem struct {
	ID             int64      `gorm:"primaryKey;column:id" json:"id"`
	MachineID      *int64     `gorm:"column:machine_id" json:"machine_id"`
	MachineRegionID *int64    `gorm:"column:machine_region_id" json:"machine_region_id"`
	DateStart      *string `gorm:"column:date_start" json:"date_start"`
	DateEnd        *string `gorm:"column:date_end" json:"date_end"`
	NoDateEnd      *bool      `gorm:"column:no_date_end" json:"no_date_end"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (MachineRegionMappingItem) TableName() string {
	return "machine_region_mapping_items"
}