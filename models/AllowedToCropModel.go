package models

import "time"

type AllowedToCropModel struct {
	ID             int64      `json:"id" gorm:"primaryKey"`
	CropID         int64      `json:"crop_id"`
	ApplicableID   int64      `json:"applicable_id"`
	ApplicableType string     `json:"applicable_type"`
	ExternalID     *string    `json:"external_id"`
	AdditionalInfo *string    `json:"additional_info"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
