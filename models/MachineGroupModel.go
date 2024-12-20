package models

import (
	"time"
)

type MachineGroup struct {
	ID             int64      `gorm:"primaryKey;column:id"`
	Name           *string    `gorm:"column:name;size:255"`
	AdditionalInfo *string    `gorm:"column:additional_info;type:text"`
	Description    *string    `gorm:"column:description;type:text"`
	ExternalID     *string    `gorm:"column:external_id;size:255"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"`
}

func (MachineGroup) TableName() string {
	return "machine_groups"
}