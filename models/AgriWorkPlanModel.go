package models

import (
	"encoding/json"
	"time"
)

type AgriWorkPlanModel struct {
	ID                          int64     `json:"id" gorm:"primaryKey"`
	Status                      string    `json:"status"`
	WorkType                    string    `json:"work_type"`
	WorkSubtype                 string    `json:"work_subtype"`
	WorkTypeID                  int64     `json:"work_type_id"`
	Season                      int       `json:"season"`
	PlannedStartDate            string    `json:"planned_start_date"`
	PlannedEndDate              string    `json:"planned_end_date"`
	AdditionalInfo              string    `json:"additional_info"`
	Description                 string    `json:"description"`
	PlannedWaterRate            float64   `json:"planned_water_rate"`
	PlannedRowSpacing           *float64  `json:"planned_row_spacing"`
	PlannedPlantSpacing         *float64  `json:"planned_plant_spacing"`
	PlannedDepth                *float64  `json:"planned_depth"`
	PlannedSpeed                *float64  `json:"planned_speed"`
	ResponsiblePersonID         *int64    `json:"responsible_person_id"`
	ExternalID                  json.RawMessage    `json:"external_id"`
	GroupableID                 int64     `json:"groupable_id"`
	GroupableType               string    `json:"groupable_type"`
	AgroRecommendationID        *int64    `json:"agro_recommendation_id"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	NotifyResponsibleUsers      bool      `json:"notify_responsible_users"`
	CurrentSeasonID             int64     `json:"current_season_id"`
	ResponsibleUserMappingItems json.RawMessage    `json:"responsible_user_mapping_items"`
}
