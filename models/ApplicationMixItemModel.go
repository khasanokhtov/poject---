package models

import (
	"time"
)

type ApplicationMixItem struct {
	ID                int            `gorm:"primaryKey;column:id" json:"id"`
	AgroOperationID   *int           `gorm:"column:agro_operation_id" json:"agro_operation_id"`
	ApplicableID      *int           `gorm:"column:applicable_id" json:"applicable_id"`
	ApplicableType    *string        `gorm:"column:applicable_type;size:255" json:"applicable_type"`
	PlannedValue      *float64       `gorm:"column:planned_value" json:"planned_value"`
	Value             *float64       `gorm:"column:value" json:"value"`
	PlannedRate       *float64       `gorm:"column:planned_rate" json:"planned_rate"`
	FactRate          *float64       `gorm:"column:fact_rate" json:"fact_rate"`
	PlannedAmount     *float64       `gorm:"column:planned_amount" json:"planned_amount"`
	FactAmount        *float64       `gorm:"column:fact_amount" json:"fact_amount"`
	ExternalID        *string        `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt         *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	IdempotencyKey    *string        `gorm:"column:idempotency_key;size:255" json:"idempotency_key"`
	CustomFactArea    *float64       `gorm:"column:custom_fact_area" json:"custom_fact_area"`
	UseCustomFactArea *bool          `gorm:"column:use_custom_fact_area" json:"use_custom_fact_area"`
}

func (ApplicationMixItem) TableName() string {
	return "application_mix_items"
}