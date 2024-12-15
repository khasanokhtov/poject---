package models

import (
	"time"
)

type GPSLoggerMappingItemModel struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	GPSLoggerID   uint      `json:"gps_logger_id" gorm:"not null"`
	MappableID    uint      `json:"mappable_id"`
	MappableType  string    `json:"mappable_type" gorm:"not null"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	NoDateEnd     bool      `json:"no_date_end"`
	ExternalID    *string   `json:"external_id"`
	AdditionalInfo string   `json:"additional_info"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
