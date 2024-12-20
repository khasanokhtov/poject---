package models

import (
	"time"
)

type GroupFolder struct {
	ID         int64      `gorm:"primaryKey;column:id"`
	ParentID   *int64     `gorm:"column:parent_id"`
	Name       *string    `gorm:"column:name;size:255"`
	ExternalID *string    `gorm:"column:external_id;size:255"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

func (GroupFolder) TableName() string {
	return "group_folders"
}