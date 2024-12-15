package models

import (
	"time"
)

type GroupFolderModel struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ParentID   *uint     `json:"parent_id"` // Может быть null
	Name       string    `json:"name" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ExternalID *string   `json:"external_id"` // Может быть null
}
