package models

import (
	"encoding/json"
	"time"
)

type AdditionalObjectModel struct {
	ID                        uint      `gorm:"primaryKey"`
	FieldGroupID              uint      `gorm:"column:field_group_id"`
	Name                      string    `gorm:"column:name"`
	ObjectType                string    `gorm:"column:object_type"`
	CalculatedArea            float64   `gorm:"column:calculated_area"`
	AdditionalInfo            string    `gorm:"column:additional_info"`
	Description               string    `gorm:"column:description"`
	GeoJSON                   string    `gorm:"column:geo_json"`
	GeometryType              string    `gorm:"column:geometry_type"`
	PointLon                  float64   `gorm:"column:point_lon"`
	PointLat                  float64   `gorm:"column:point_lat"`
	AdministrativeAreaName    string    `gorm:"column:administrative_area_name"`
	SubAdministrativeAreaName string    `gorm:"column:subadministrative_area_name"`
	Locality                  string    `gorm:"column:locality"`
	Additional                json.RawMessage    `gorm:"type:jsonb;column:additional"`
	ExternalID                *string   `gorm:"column:external_id"`
	CreatedAt                 time.Time `gorm:"column:created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at"`
}