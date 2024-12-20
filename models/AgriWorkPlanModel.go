package models

import (
	"time"
)

type AgriWorkPlan struct {
	ID                        int            `gorm:"primaryKey;column:id"`
	Status                    *string        `gorm:"column:status;size:255"`
	WorkType                  *string        `gorm:"column:work_type;size:255"`
	WorkSubtype               *string        `gorm:"column:work_subtype;size:255"`
	WorkTypeID                *int           `gorm:"column:work_type_id"`
	Season                    *int           `gorm:"column:season"`
	PlannedStartDate          *time.Time     `gorm:"column:planned_start_date"`
	PlannedEndDate            *time.Time     `gorm:"column:planned_end_date"`
	AdditionalInfo            *string        `gorm:"column:additional_info;size:255"`
	Description               *string        `gorm:"column:description;size:255"`
	PlannedWaterRate          *float64       `gorm:"column:planned_water_rate"`
	PlannedRowSpacing         *float64       `gorm:"column:planned_row_spacing"`
	PlannedPlantSpacing       *float64       `gorm:"column:planned_plant_spacing"`
	PlannedDepth              *float64       `gorm:"column:planned_depth"`
	PlannedSpeed              *float64       `gorm:"column:planned_speed"`
	ResponsiblePersonID       *int           `gorm:"column:responsible_person_id"`
	ExternalID                *string        `gorm:"column:external_id;size:255"`
	GroupableID               *int           `gorm:"column:groupable_id"`
	GroupableType             *string        `gorm:"column:groupable_type;size:255"`
	AgroRecommendationID      *int           `gorm:"column:agro_recommendation_id"`
	CreatedAt                 *time.Time     `gorm:"column:created_at"`
	UpdatedAt                 *time.Time     `gorm:"column:updated_at"`
	ResponsibleUserIDs        *string        `gorm:"column:responsible_user_ids;type:jsonb"`
	NotifyResponsibleUsers    *bool          `gorm:"column:notify_responsible_users;default:true"`
	CurrentSeasonID           *int           `gorm:"column:current_season_id"`
}

func (AgriWorkPlan) TableName() string {
	return "agri_work_plans"
}
