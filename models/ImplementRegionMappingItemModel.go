package models

import (
	"time"
)

// ImplementRegionMappingItemModel - структура для хранения данных о привязке региона к навесному оборудованию
type ImplementRegionMappingItemModel struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	ImplementID     uint       `json:"implement_id"`
	MachineRegionID uint       `json:"machine_region_id"`
	DateStart       CustomDate `json:"date_start"`
	DateEnd         *CustomDate `json:"date_end"`
	NoDateEnd       bool       `json:"no_date_end"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
