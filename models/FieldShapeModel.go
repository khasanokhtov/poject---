package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// CustomString - кастомный тип для обработки строк или чисел
type CustomString string

func (cs *CustomString) UnmarshalJSON(data []byte) error {
	// Попробуем декодировать как строку
	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		*cs = CustomString(asString)
		return nil
	}

	// Если декодирование строки не удалось, пробуем декодировать как число
	var asNumber float64
	if err := json.Unmarshal(data, &asNumber); err == nil {
		*cs = CustomString(fmt.Sprintf("%g", asNumber)) // Конвертируем число в строку
		return nil
	}

	return fmt.Errorf("не удалось декодировать CustomString: %s", string(data))
}

type FieldShapeModel struct {
	ID                    uint        `json:"id" gorm:"primaryKey"`
	FieldID               uint        `json:"field_id"`
	StartTime             time.Time   `json:"start_time"`
	EndTime               time.Time   `json:"end_time"`
	CalculatedArea        float64     `json:"calculated_area"`
	LegalArea             float64     `json:"legal_area"`
	TillableArea          float64     `json:"tillable_area"`
	SimplifiedShape       string      `json:"simplified_shape" gorm:"type:text"`
	ShapeSimplifiedGeoJSON string     `json:"shape_simplified_geojson" gorm:"type:text"`
	PointLon              float64     `json:"point_lon"`
	PointLat              float64     `json:"point_lat"`
	ExternalID            *string     `json:"external_id"`
	LegalPerimeter        CustomString `json:"legal_perimeter"` // Используем кастомный тип
	Description           *string     `json:"description"`
	LandQualityAssessment *string     `json:"land_quality_assessment"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
}
