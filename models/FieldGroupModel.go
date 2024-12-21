package models

import (
	"time"
)

type FieldGroup struct {
	ID                          int64      `gorm:"primaryKey;column:id" json:"id"`
	GroupFolderID               *int64     `gorm:"column:group_folder_id" json:"group_folder_id"`
	Name                        *string    `gorm:"column:name;size:255" json:"name"`
	Description                 *string    `gorm:"column:description;type:text" json:"description"`
	AdministrativeAreaName      *string    `gorm:"column:administrative_area_name;size:255" json:"administrative_area_name"`
	SubAdministrativeAreaName   *string    `gorm:"column:subadministrative_area_name;size:255" json:"subadministrative_area_name"`
	Locality                    *string    `gorm:"column:locality;size:255" json:"locality"`
	Hidden                      *bool      `gorm:"column:hidden;default:false" json:"hidden"`
	ExternalID                  *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt                   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	IdempotencyKey              *string    `gorm:"column:idempotency_key;size:255" json:"idempotency_key"`
	LegalEntity                 *string    `gorm:"column:legal_entity;size:255" json:"legal_entity"`
	MachineTaskDefaultDuration  *int64     `gorm:"column:machine_task_default_duration" json:"machine_task_default_duration"`
	AccountingPeriodClosingDate *string    `gorm:"column:accounting_period_closing_date" json:"accounting_period_closing_date"`
	MachineTaskDefaultStartTime *string    `gorm:"column:machine_task_default_start_time;size:255" json:"machine_task_default_start_time"`
}

func (FieldGroup) TableName() string {
	return "field_groups"
}