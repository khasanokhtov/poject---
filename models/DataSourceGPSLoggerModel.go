package models

import (
	"time"
)

type DataSourceGPSLoggerModel struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	MappableID     uint      `json:"mappable_id"`
	MappableType   string    `json:"mappable_type" gorm:"not null"`
	GPSLoggerID    uint      `json:"gps_logger_id"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	AdditionalInfo *string   `json:"additional_info"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}