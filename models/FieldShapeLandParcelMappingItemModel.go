package models

import (
	"time"
)

type FieldShapeLandParcelMappingItemModel struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	FieldShapeID  uint      `json:"field_shape_id"`
	LandParcelID  uint      `json:"land_parcel_id"`
	ExternalID    *string   `json:"external_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
