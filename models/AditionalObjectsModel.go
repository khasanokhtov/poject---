package models

type AdditionalObject struct {
	ID                        uint     `gorm:"primaryKey;column:id" json:"id"`
	FieldGroupID              *uint    `gorm:"column:field_group_id" json:"field_group_id"`
	Name                      *string  `gorm:"column:name;size:255" json:"name"`
	ObjectType                *string  `gorm:"column:object_type;size:255" json:"object_type"`
	CalculatedArea            *float64 `gorm:"column:calculated_area" json:"calculated_area"`
	AdditionalInfo            *string  `gorm:"column:additional_info;size:255" json:"additional_info"`
	Description               *string  `gorm:"column:description;size:255" json:"description"`
	GeoJSON                   *string  `gorm:"column:geo_json;type:jsonb" json:"geo_json"`
	GeometryType              *string  `gorm:"column:geometry_type;size:255" json:"geometry_type"`
	PointLon                  *float64 `gorm:"column:point_lon" json:"point_lon"`
	PointLat                  *float64 `gorm:"column:point_lat" json:"point_lat"`
	AdministrativeAreaName    *string  `gorm:"column:administrative_area_name;size:255" json:"administrative_area_name"`
	SubadministrativeAreaName *string  `gorm:"column:subadministrative_area_name;size:255" json:"subadministrative_area_name"`
	Locality                  *string  `gorm:"column:locality;size:255" json:"locality"`
	ExternalID                *string  `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt                 *string  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 *string  `gorm:"column:updated_at" json:"updated_at"`
}

func (AdditionalObject) TableName() string {
	return "additional_object"
}