package models

import (
	"time"
)

type WorkTypeGroup struct {
	ID          uint       `gorm:"primaryKey;column:id" json:"id"`
	Name        *string    `gorm:"column:name;size:255" json:"name"`
	StandardName *string   `gorm:"column:standard_name;size:255" json:"standard_name"`
	Description *string    `gorm:"column:description;size:255" json:"description"`
	ExternalID  *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (WorkTypeGroup) TableName() string {
	return "work_type_groups"
}
