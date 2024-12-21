package models

import (
	"time"
)

type MachineGroup struct {
	ID             int64      `gorm:"primaryKey;column:id" json:"id"`
	Name           *string    `gorm:"column:name;size:255" json:"name"`
	AdditionalInfo *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	Description    *string    `gorm:"column:description;type:text" json:"description"`
	ExternalID     *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (MachineGroup) TableName() string {
	return "machine_groups"
}