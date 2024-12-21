package models

import (
	"time"
)

type Field struct {
	ID                         int64      `gorm:"primaryKey;column:id" json:"id"`
	FieldGroupID               *int64     `gorm:"column:field_group_id" json:"field_group_id"`
	Name                       *string    `gorm:"column:name;size:255" json:"name"`
	AdditionalName             *string    `gorm:"column:additional_name;size:255" json:"additional_name"`
	Description                *string    `gorm:"column:description;type:text" json:"description"`
	MoistureZone               *string    `gorm:"column:moisture_zone;size:255" json:"moisture_zone"`
	ShapeSimplifiedGeojson     *string    `gorm:"column:shape_simplified_geojson;type:text" json:"shape_simplified_geojson"`
	CurrentShapeID             *int64     `gorm:"column:current_shape_id" json:"current_shape_id"`
	CalculatedArea             *float64   `gorm:"column:calculated_area" json:"calculated_area"`
	LegalArea                  *float64   `gorm:"column:legal_area" json:"legal_area"`
	TillableArea               *float64   `gorm:"column:tillable_area" json:"tillable_area"`
	AdministrativeAreaName     *string    `gorm:"column:administrative_area_name;size:255" json:"administrative_area_name"`
	SubAdministrativeAreaName  *string    `gorm:"column:subadministrative_area_name;size:255" json:"subadministrative_area_name"`
	Locality                   *string    `gorm:"column:locality;size:255" json:"locality"`
	FieldType                  *string    `gorm:"column:field_type;size:255" json:"field_type"`
	ParentID                   *int64     `gorm:"column:parent_id" json:"parent_id"`
	Lat                        *float64   `gorm:"column:lat" json:"lat"`
	Long                       *float64   `gorm:"column:long" json:"long"`
	AdditionalInfo             *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	ExternalID                 *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt                  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Field) TableName() string {
	return "fields"
}