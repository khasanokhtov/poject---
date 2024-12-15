package models

import (
	"time"
)

type FieldModel struct {
	ID                         uint        `json:"id" gorm:"primaryKey"`
	FieldGroupID               uint        `json:"field_group_id"`
	Name                       string      `json:"name" gorm:"not null"`
	AdditionalName             string      `json:"additional_name"`
	AdditionalInfo             string      `json:"additional_info"`
	Description                string      `json:"description"`
	AdministrativeAreaName     string      `json:"administrative_area_name"`
	SubAdministrativeAreaName  string      `json:"subadministrative_area_name"`
	Locality                   string      `json:"locality"`
	IdempotencyKey             *string     `json:"idempotency_key"`
	RegionID                   *uint       `json:"region_id"`
	CountryID                  uint        `json:"country_id"`
	DistrictID                 *uint       `json:"district_id"`
	FieldType                  string      `json:"field_type"`
	ParentID                   *uint       `json:"parent_id"`
	Lat                        float64     `json:"lat"`
	Long                       float64     `json:"long"`
	ExternalID                 *string     `json:"external_id"`
	CreatedAt                  time.Time   `json:"created_at"`
	UpdatedAt                  time.Time   `json:"updated_at"`
	XCustomFieldsData          JSONB       `json:"x_custom_fields_data" gorm:"type:jsonb"` // Обработка JSONB
}