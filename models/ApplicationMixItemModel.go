package models

import "time"

type ApplicationMixItemModel struct {
	ID                int64      `json:"id" gorm:"primaryKey"`
	AgroOperationID   int64      `json:"agro_operation_id"`
	ApplicableID      int64      `json:"applicable_id"`
	ApplicableType    string     `json:"applicable_type"`
	PlannedValue      float64    `json:"planned_value"`
	Value             float64    `json:"value"`
	PlannedRate       float64    `json:"planned_rate"`
	FactRate          float64    `json:"fact_rate"`
	PlannedAmount     float64    `json:"planned_amount"`
	FactAmount        float64    `json:"fact_amount"`
	ExternalID        *string    `json:"external_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	IdempotencyKey    *string    `json:"idempotency_key"`
	CustomFactArea    float64    `json:"custom_fact_area"`
	UseCustomFactArea bool       `json:"use_custom_fact_area"`
}
