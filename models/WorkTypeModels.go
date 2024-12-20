package models

import (
	"time"
)

type WorkType struct {
	ID                           uint       `gorm:"primaryKey;column:id"`
	WorkTypeGroupID              *uint      `gorm:"column:work_type_group_id"`
	Name                         *string    `gorm:"column:name;size:255"`
	Agri                         *bool      `gorm:"column:agri"`
	Application                  *bool      `gorm:"column:application"`
	Sowing                       *bool      `gorm:"column:sowing"`
	ReSowing                     *bool      `gorm:"column:re_sowing"`
	AdditionalSowing             *bool      `gorm:"column:additional_sowing"`
	Harvesting                   *bool      `gorm:"column:harvesting"`
	Soil                         *bool      `gorm:"column:soil"`
	StandardName                 *string    `gorm:"column:standard_name;size:255"`
	Hidden                       *bool      `gorm:"column:hidden"`
	Description                  *string    `gorm:"column:description;size:255"`
	ExternalID                   *string    `gorm:"column:external_id;size:255"`
	DisableFirstTrackRuleCoverage *bool     `gorm:"column:disable_first_track_rule_coverage"`
	CreatedAt                    *time.Time `gorm:"column:created_at"`
	UpdatedAt                    *time.Time `gorm:"column:updated_at"`
}

func (WorkType) TableName() string {
	return "work_types"
}
