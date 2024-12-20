package models

import (
	"time"
)

type Chemical struct {
	ID                           int64     `gorm:"primaryKey;column:id"`
	Name                         string    `gorm:"column:name;size:255"`
	ChemicalType                 *string   `gorm:"column:chemical_type;size:255"`
	UnitsOfMeasurement           *string   `gorm:"column:units_of_measurement;size:255"`
	AdditionalInfo               *string   `gorm:"column:additional_info;type:text"`
	Description                  *string   `gorm:"column:description;type:text"`
	CreatedAt                    *time.Time `gorm:"column:created_at"`
	UpdatedAt                    *time.Time `gorm:"column:updated_at"`
	Archived                     *bool     `gorm:"column:archived;default:false"`
	ExternalID                   *string   `gorm:"column:external_id;size:255"`
	ToxicityClass                *int      `gorm:"column:toxicity_class"`
	ActionTerm                   *int      `gorm:"column:action_term"`
	ActionTermUnits              *string   `gorm:"column:action_term_units;size:255"`
	ActiveSubstance              *string   `gorm:"column:active_substance;size:255"`
	DrugForm                     *string   `gorm:"column:drug_form;size:255"`
	InfluenceMethod              *string   `gorm:"column:influence_method;size:255"`
	BeesIsolatingRecommendedTerm *int      `gorm:"column:bees_isolating_recommended_term"`
	BeesIsolatingRecommendedTermUnits *string `gorm:"column:bees_isolating_recommended_term_units;size:255"`
	SaleTerm                     *string   `gorm:"column:sale_term;size:255"`
	TermOfUse                    *string   `gorm:"column:term_of_use;size:255"`
	WhItemID                     *int64    `gorm:"column:wh_item_id"`
	WhItemBaseUnitID             *int64    `gorm:"column:wh_item_base_unit_id"`
	ActiveSubstances             *string   `gorm:"column:active_substances;type:jsonb"`
}

func (Chemical) TableName() string {
	return "chemicals"
}