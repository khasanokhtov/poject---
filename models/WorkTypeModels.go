package models

import (
	"time"
)

type WorkType struct {
	ID                           uint       `gorm:"primaryKey;column:id" json:"id"`
	WorkTypeGroupID              *uint      `gorm:"column:work_type_group_id" json:"work_type_group_id"`
	Name                         *string    `gorm:"column:name;size:255" json:"name"`
	Agri                         *bool      `gorm:"column:agri" json:"agri"`
	Application                  *bool      `gorm:"column:application" json:"application"`
	Sowing                       *bool      `gorm:"column:sowing" json:"sowing"`
	ReSowing                     *bool      `gorm:"column:re_sowing" json:"re_sowing"`
	AdditionalSowing             *bool      `gorm:"column:additional_sowing" json:"additional_sowing"`
	Harvesting                   *bool      `gorm:"column:harvesting" json:"harvesting"`
	Soil                         *bool      `gorm:"column:soil" json:"soil"`
	StandardName                 *string    `gorm:"column:standard_name;size:255" json:"standard_name"`
	Hidden                       *bool      `gorm:"column:hidden" json:"hidden"`
	Description                  *string    `gorm:"column:description;size:255" json:"description"`
	ExternalID                   *string    `gorm:"column:external_id;size:255" json:"external_id"`
	DisableFirstTrackRuleCoverage *bool     `gorm:"column:disable_first_track_rule_coverage" json:"disable_first_track_rule_coverage"`
	CreatedAt                    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (WorkType) TableName() string {
	return "work_types"
}
