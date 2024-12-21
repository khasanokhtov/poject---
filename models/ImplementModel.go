package models

import (
	"time"
)

type Implement struct {
	ID                 int64      `gorm:"primaryKey;column:id" json:"id"`
	Name               *string    `gorm:"column:name;size:255" json:"name"`
	Model              *string    `gorm:"column:model;size:255" json:"model"`
	Manufacturer       *string    `gorm:"column:manufacturer;size:255" json:"manufacturer"`
	Year               *int64     `gorm:"column:year" json:"year"`
	RegistrationNumber *string    `gorm:"column:registration_number;size:255" json:"registration_number"`
	InventoryNumber    *string    `gorm:"column:inventory_number;size:255" json:"inventory_number"`
	ImplementType      *string    `gorm:"column:implement_type;size:255" json:"implement_type"`
	Width              *float64   `gorm:"column:width" json:"width"`
	OfficialWidth      *float64   `gorm:"column:official_width" json:"official_width"`
	AvatarID           *int64     `gorm:"column:avatar_id" json:"avatar_id"`
	ChassisSerialNumber *string   `gorm:"column:chassis_serial_number;size:255" json:"chassis_serial_number"`
	LegalCompany       *string    `gorm:"column:legal_company;size:255" json:"legal_company"`
	Description        *string    `gorm:"column:description;type:text" json:"description"`
	AdditionalInfo     *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	VariableWidth      *bool      `gorm:"column:variable_width" json:"variable_width"`
	MinWidth           *float64   `gorm:"column:min_width" json:"min_width"`
	MaxWidth           *float64   `gorm:"column:max_width" json:"max_width"`
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`
	ExternalID         *string    `gorm:"column:external_id;size:255" json:"external_id"`
	Virtual            *bool      `gorm:"column:virtual" json:"virtual"`
}

func (Implement) TableName() string {
	return "implements"
}