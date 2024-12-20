package models

import (
	"time"
)

type ApplicationMixItem struct {
	ID                int            `gorm:"primaryKey;column:id"`
	AgroOperationID   *int           `gorm:"column:agro_operation_id"`
	ApplicableID      *int           `gorm:"column:applicable_id"`
	ApplicableType    *string        `gorm:"column:applicable_type;size:255"`
	PlannedValue      *float64       `gorm:"column:planned_value"`
	Value             *float64       `gorm:"column:value"`
	PlannedRate       *float64       `gorm:"column:planned_rate"`
	FactRate          *float64       `gorm:"column:fact_rate"`
	PlannedAmount     *float64       `gorm:"column:planned_amount"`
	FactAmount        *float64       `gorm:"column:fact_amount"`
	ExternalID        *string        `gorm:"column:external_id;size:255"`
	CreatedAt         *time.Time     `gorm:"column:created_at"`
	UpdatedAt         *time.Time     `gorm:"column:updated_at"`
	IdempotencyKey    *string        `gorm:"column:idempotency_key;size:255"`
}

func (ApplicationMixItem) TableName() string {
	return "application_mix_items"
}