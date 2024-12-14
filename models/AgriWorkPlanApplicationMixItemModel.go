package models

import (
	"time"
)

type AgriWorkPlanApplicationMixItemModel struct {
	ID              int64     `json:"id" gorm:"primaryKey"`
	ApplicableID    int64     `json:"applicable_id"`
	ApplicableType  string    `json:"applicable_type"`
	Amount          float64   `json:"amount"`
	AdditionalInfo  *string   `json:"additional_info"`
	Rate            float64   `json:"rate"`
	AgriWorkPlanID  int64     `json:"agri_work_plan_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
