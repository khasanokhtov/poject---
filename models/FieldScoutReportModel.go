package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type ImageDetails struct {
	URL         string   `json:"url"`
	Lat         *float64 `json:"lat"`
	Lon         *float64 `json:"lon"`
	MD5         *string  `json:"md5"`
	Preview200  string   `json:"preview_200"`
	Preview400  string   `json:"preview_400"`
	Preview1000 string   `json:"preview_1000"`
}

// NDVI - кастомный тип для обработки строк и чисел
type NDVI float64

func (n *NDVI) UnmarshalJSON(data []byte) error {
	// Попробуем декодировать как число
	var asFloat float64
	if err := json.Unmarshal(data, &asFloat); err == nil {
		*n = NDVI(asFloat)
		return nil
	}

	// Если декодирование числа не удалось, пробуем декодировать как строку
	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		// Проверяем, не является ли строка пустой
		if asString == "" {
			*n = NDVI(0) // Устанавливаем значение по умолчанию (например, 0)
			return nil
		}

		parsed, err := strconv.ParseFloat(asString, 64)
		if err != nil {
			return fmt.Errorf("ошибка парсинга NDVI из строки: %w", err)
		}
		*n = NDVI(parsed)
		return nil
	}

	return fmt.Errorf("не удалось декодировать NDVI: %s", string(data))
}

// MarshalJSON - реализация для возврата значения как float64
func (n NDVI) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(n))
}


type FieldScoutReportModel struct {
	ID                  uint        `json:"id" gorm:"primaryKey"`
	FieldID             uint        `json:"field_id"`
	UserID              uint        `json:"user_id"`
	ReportTime          time.Time   `json:"report_time"`
	Season              uint        `json:"season"`
	GrowthScale         *string     `json:"growth_scale"`
	GrowthStage         *string     `json:"growth_stage"`
	AdditionalInfo      *string     `json:"additional_info"`
	GrowthStageID       *uint       `json:"growth_stage_id"`
	FieldShapeID        uint        `json:"field_shape_id"`
	Image1              ImageDetails `gorm:"embedded;embeddedPrefix:image1_"`
	Image2              ImageDetails `gorm:"embedded;embeddedPrefix:image2_"`
	Image3              ImageDetails `gorm:"embedded;embeddedPrefix:image3_"`
	EarsCount           *uint       `json:"ears_count"`
	PlantsCount         *uint       `json:"plants_count"`
	GroundCover         *float64    `json:"ground_cover"`
	CreatedByUserAt     time.Time   `json:"created_by_user_at"`
	UpdatedByUserAt     time.Time   `json:"updated_by_user_at"`
	ExternalID          *string     `json:"external_id"`
	ScoutingTaskID      *uint       `json:"scouting_task_id"`
	ScoutReportTemplateID *uint     `json:"scout_report_template_id"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	RiskYieldDecreasing bool        `json:"risk_yield_decreasing"`
	IdempotencyKey      *string     `json:"idempotency_key"`
	NotifyManagement    bool        `json:"notify_management"`
	HistoryItemID       uint        `json:"history_item_id"`
	FieldNDVI           NDVI        `json:"field_ndvi"` // Используем кастомный тип NDVI
}
