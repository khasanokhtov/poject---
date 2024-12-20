package models

import (
	"time"
)

type Field struct {
	ID                         int64      `gorm:"primaryKey;column:id"`
	FieldGroupID               *int64     `gorm:"column:field_group_id"`
	Name                       *string    `gorm:"column:name;size:255"`
	AdditionalName             *string    `gorm:"column:additional_name;size:255"`
	Description                *string    `gorm:"column:description;type:text"`
	MoistureZone               *string    `gorm:"column:moisture_zone;size:255"`
	ShapeSimplifiedGeojson     *string    `gorm:"column:shape_simplified_geojson;type:text"`
	CurrentShapeID             *int64     `gorm:"column:current_shape_id"`
	CalculatedArea             *float64   `gorm:"column:calculated_area"`
	LegalArea                  *float64   `gorm:"column:legal_area"`
	TillableArea               *float64   `gorm:"column:tillable_area"`
	AdministrativeAreaName     *string    `gorm:"column:administrative_area_name;size:255"`
	SubAdministrativeAreaName  *string    `gorm:"column:subadministrative_area_name;size:255"`
	Locality                   *string    `gorm:"column:locality;size:255"`
	FieldType                  *string    `gorm:"column:field_type;size:255"`
	ParentID                   *int64     `gorm:"column:parent_id"`
	Lat                        *float64   `gorm:"column:lat"`
	Long                       *float64   `gorm:"column:long"`
	AdditionalInfo             *string    `gorm:"column:additional_info;type:text"`
	ExternalID                 *string    `gorm:"column:external_id;size:255"`
	CreatedAt                  *time.Time `gorm:"column:created_at"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at"`
}

func (Field) TableName() string {
	return "fields"
}