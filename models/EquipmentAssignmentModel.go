package models

import "time"

type EquipmentAssignmentModel struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	EquippableType     string    `json:"equippable_type" gorm:"not null"`
	EquippableID       uint      `json:"equippable_id" gorm:"not null"`
	EquipmentHolderType string   `json:"equipment_holder_type" gorm:"not null"`
	EquipmentHolderID  uint      `json:"equipment_holder_id" gorm:"not null"`
	StartTime          time.Time `json:"start_time"`
	EndTime            time.Time `json:"end_time"`
	NoEndTime          bool      `json:"no_end_time"`
	ExternalID         *string   `json:"external_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
