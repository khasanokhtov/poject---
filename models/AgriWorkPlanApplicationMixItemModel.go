package models

import (
	"time"
)

type AgriWorkPlanApplicationMixItem struct {
	ID             int64     `gorm:"primaryKey;column:id" json:"id"`
	ApplicableID   *int64    `gorm:"column:applicable_id" json:"applicable_id"`
	ApplicableType *string   `gorm:"column:applicable_type;size:255" json:"applicable_type"`
	Amount         *float64  `gorm:"column:amount" json:"amount"`
	AdditionalInfo *string   `gorm:"column:additional_info;size:255" json:"additional_info"`
	Rate           *float64  `gorm:"column:rate" json:"rate"`
	AgriWorkPlanID *int64    `gorm:"column:agri_work_plan_id" json:"agri_work_plan_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (AgriWorkPlanApplicationMixItem) TableName() string {
	return "agri_work_plan_application_mix_ites"
}