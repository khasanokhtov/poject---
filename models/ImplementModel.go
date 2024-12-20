package models

import (
	"time"
)

type Implement struct {
	ID                 int64      `gorm:"primaryKey;column:id"`
	Name               *string    `gorm:"column:name;size:255"`
	Model              *string    `gorm:"column:model;size:255"`
	Manufacturer       *string    `gorm:"column:manufacturer;size:255"`
	Year               *int64     `gorm:"column:year"`
	RegistrationNumber *string    `gorm:"column:registration_number;size:255"`
	InventoryNumber    *string    `gorm:"column:inventory_number;size:255"`
	ImplementType      *string    `gorm:"column:implement_type;size:255"`
	Width              *float64   `gorm:"column:width"`
	OfficialWidth      *float64   `gorm:"column:official_width"`
	AvatarID           *int64     `gorm:"column:avatar_id"`
	ChassisSerialNumber *string   `gorm:"column:chassis_serial_number;size:255"`
	LegalCompany       *string    `gorm:"column:legal_company;size:255"`
	Description        *string    `gorm:"column:description;type:text"`
	AdditionalInfo     *string    `gorm:"column:additional_info;type:text"`
	VariableWidth      *bool      `gorm:"column:variable_width"`
	MinWidth           *float64   `gorm:"column:min_width"`
	MaxWidth           *float64   `gorm:"column:max_width"`
	CreatedAt          *time.Time `gorm:"column:created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at"`
	ExternalID         *string    `gorm:"column:external_id;size:255"`
	Virtual            *bool      `gorm:"column:virtual"`
}

func (Implement) TableName() string {
	return "implements"
}