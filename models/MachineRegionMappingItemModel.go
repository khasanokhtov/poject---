package models

import (
	"time"
)

type MachineRegionMappingItem struct {
	ID             int64      `gorm:"primaryKey;column:id"`
	MachineID      *int64     `gorm:"column:machine_id"`
	MachineRegionID *int64    `gorm:"column:machine_region_id"`
	DateStart      *time.Time `gorm:"column:date_start"`
	DateEnd        *time.Time `gorm:"column:date_end"`
	NoDateEnd      *bool      `gorm:"column:no_date_end"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"`
}

func (MachineRegionMappingItem) TableName() string {
	return "machine_region_mapping_items"
}