package models

import (
	"time"
)

type GPSLoggerModel struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	LoggerType          string    `json:"logger_type" gorm:"not null"`
	IMEI                string    `json:"imei" gorm:"not null"`
	PhoneNumber         *string   `json:"phone_number"`
	SerialNumber        string    `json:"serial_number"`
	Description         *string   `json:"description"`
	MaxTimeBetweenPackets *uint   `json:"max_time_between_packets"`
	ExternalID          *string   `json:"external_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
