package models

import (
	"time"
)

type MachineRegion struct {
	ID            int64      `gorm:"primaryKey;column:id"  json:"id"`
	Name          *string    `gorm:"column:name;size:255"  json:"name"`
	Ancestry      *string    `gorm:"column:ancestry;size:255"  json:"ancestry"`
	Description   *string    `gorm:"column:description;type:text"  json:"description"`
	AdditionalInfo *string   `gorm:"column:additional_info;type:text"  json:"additional_info"`
	CreatedAt     *time.Time `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"  json:"updated_at"`
	ExternalID    *string    `gorm:"column:external_id;size:255"  json:"external_id"`
}

func (MachineRegion) TableName() string {
	return "machine_region"
}