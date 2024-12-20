package models

import (
	"time"
)

type MachineRegion struct {
	ID            int64      `gorm:"primaryKey;column:id"`
	Name          *string    `gorm:"column:name;size:255"`
	Ancestry      *string    `gorm:"column:ancestry;size:255"`
	Description   *string    `gorm:"column:description;type:text"`
	AdditionalInfo *string   `gorm:"column:additional_info;type:text"`
	CreatedAt     *time.Time `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
	ExternalID    *string    `gorm:"column:external_id;size:255"`
}

func (MachineRegion) TableName() string {
	return "machine_region"
}