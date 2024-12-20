package models

import (
	"time"
)

type FieldGroup struct {
	ID                          int64      `gorm:"primaryKey;column:id"`
	GroupFolderID               *int64     `gorm:"column:group_folder_id"`
	Name                        *string    `gorm:"column:name;size:255"`
	Description                 *string    `gorm:"column:description;type:text"`
	AdministrativeAreaName      *string    `gorm:"column:administrative_area_name;size:255"`
	SubAdministrativeAreaName   *string    `gorm:"column:subadministrative_area_name;size:255"`
	Locality                    *string    `gorm:"column:locality;size:255"`
	Hidden                      *bool      `gorm:"column:hidden;default:false"`
	ExternalID                  *string    `gorm:"column:external_id;size:255"`
	CreatedAt                   *time.Time `gorm:"column:created_at"`
	UpdatedAt                   *time.Time `gorm:"column:updated_at"`
	IdempotencyKey              *string    `gorm:"column:idempotency_key;size:255"`
	LegalEntity                 *string    `gorm:"column:legal_entity;size:255"`
	MachineTaskDefaultDuration  *int64     `gorm:"column:machine_task_default_duration"`
	AccountingPeriodClosingDate *time.Time `gorm:"column:accounting_period_closing_date"`
	MachineTaskDefaultStartTime *string    `gorm:"column:machine_task_default_start_time;size:255"`
}

func (FieldGroup) TableName() string {
	return "field_groups"
}