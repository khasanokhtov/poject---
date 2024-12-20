package models

import (
	"time"
)

type WorkTypeGroup struct {
	ID          uint       `gorm:"primaryKey;column:id"`
	Name        *string    `gorm:"column:name;size:255"`
	StandardName *string   `gorm:"column:standard_name;size:255"`
	Description *string    `gorm:"column:description;size:255"`
	ExternalID  *string    `gorm:"column:external_id;size:255"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (WorkTypeGroup) TableName() string {
	return "work_type_groups"
}
