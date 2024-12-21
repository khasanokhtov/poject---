package models

import (
	"time"
)

type GroupFolder struct {
	ID         int64      `gorm:"primaryKey;column:id" json:"id"`
	ParentID   *int64     `gorm:"column:parent_id" json:"parent_id"`
	Name       *string    `gorm:"column:name;size:255" json:"name"`
	ExternalID *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (GroupFolder) TableName() string {
	return "group_folders"
}