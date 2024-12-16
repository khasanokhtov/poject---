package models

import "time"

// WorkType - модель для таблицы work_types
type WorkType struct {
	ID                           int        `json:"id" gorm:"primaryKey"`
	WorkTypeGroupID              int        `json:"work_type_group_id"`
	Name                         string     `json:"name"`
	Agri                         bool       `json:"agri"`
	Application                  bool       `json:"application"`
	Sowing                       bool       `json:"sowing"`
	ReSowing                     bool       `json:"re_sowing"`
	AdditionalSowing             bool       `json:"additional_sowing"`
	Harvesting                   bool       `json:"harvesting"`
	Soil                         bool       `json:"soil"`
	StandardName                 string     `json:"standard_name"`
	Hidden                       bool       `json:"hidden"`
	Description                  *string    `json:"description"` // Может быть пустым или null
	ExternalID                   *string    `json:"external_id"` // Может быть null
	DisableFirstTrackRuleCoverage bool      `json:"disable_first_track_rule_coverage"`
	CreatedAt                    time.Time  `json:"created_at"`
	UpdatedAt                    time.Time  `json:"updated_at"`
}
