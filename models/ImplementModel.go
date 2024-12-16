package models

import "time"

// ImplementModel - структура для хранения данных о навесном оборудовании
type ImplementModel struct {
	ID                 uint       `json:"id" gorm:"primaryKey"`
	Name               string     `json:"name"`
	Model              string     `json:"model"`
	Manufacturer       string     `json:"manufacturer"`
	Year               uint       `json:"year"`
	RegistrationNumber string     `json:"registration_number"`
	InventoryNumber    string     `json:"inventory_number"`
	ImplementType      string     `json:"implement_type"`
	Width              float64    `json:"width"`
	OfficialWidth      *float64   `json:"official_width"`
	AvatarID           uint       `json:"avatar_id"`
	ChassisSerialNumber string    `json:"chassis_serial_number"`
	LegalCompany       string     `json:"legal_company"`
	Description        string     `json:"description"`
	Additional         JSONB      `json:"additional" gorm:"type:jsonb"`
	AdditionalInfo     string     `json:"additional_info"`
	VariableWidth      bool       `json:"variable_width"`
	MinWidth           *float64   `json:"min_width"`
	MaxWidth           *float64   `json:"max_width"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	ExternalID         *string    `json:"external_id"`
	Virtual            bool       `json:"virtual"`
}
