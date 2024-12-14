package models

import "time"

// AgronomistAssignmentModel представляет модель данных для агрономических назначений
type AgronomistAssignmentModel struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	UserID       int64     `json:"user_id"`
	FieldGroupID int64     `json:"field_group_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}