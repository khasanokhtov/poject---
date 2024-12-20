package models

import (
	"time"
)

type AdditionalObject struct {
	ID                        int            `gorm:"primaryKey;column:id"`
	FieldGroupID              *int           `gorm:"column:field_group_id"`
	Name                      *string        `gorm:"column:name;size:255"`
	ObjectType                *string        `gorm:"column:object_type;size:255"`
	CalculatedArea            *float64       `gorm:"column:calculated_area"`
	AdditionalInfo            *string        `gorm:"column:additional_info;type:text"`
	Description               *string        `gorm:"column:description;type:text"`
	GeoJSON                   *string        `gorm:"column:geo_json;type:text"`
	GeometryType              *string        `gorm:"column:geometry_type;size:255"`
	PointLon                  *float64       `gorm:"column:point_lon"`
	PointLat                  *float64       `gorm:"column:point_lat"`
	AdministrativeAreaName    *string        `gorm:"column:administrative_area_name;size:255"`
	SubAdministrativeAreaName *string        `gorm:"column:subadministrative_area_name;size:255"`
	Locality                  *string        `gorm:"column:locality;size:255"`
	Icon                      *string        `gorm:"column:icon;type:text"`
	PointSize                 *string        `gorm:"column:point_size;type:text"`
	PointColor                *string        `gorm:"column:point_color;type:text"`
	ExternalID                *string        `gorm:"column:external_id;size:255"`
	CreatedAt                 *time.Time     `gorm:"column:created_at"`
	UpdatedAt                 *time.Time     `gorm:"column:updated_at"`
}

func (AdditionalObject) TableName() string {
	return "additional_object"
}