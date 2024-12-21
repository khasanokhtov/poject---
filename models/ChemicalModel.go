package models

import (
	"time"
)

type Chemical struct {
	ID                           int64     `gorm:"primaryKey;column:id" json:"id"`
	Name                         string    `gorm:"column:name;size:255" json:"name"`
	ChemicalType                 *string   `gorm:"column:chemical_type;size:255" json:"chemical_type"`
	UnitsOfMeasurement           *string   `gorm:"column:units_of_measurement;size:255" json:"units_of_measurement"`
	AdditionalInfo               *string   `gorm:"column:additional_info;type:text" json:"additional_info"`
	Description                  *string   `gorm:"column:description;type:text" json:"description"`
	CreatedAt                    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Archived                     *bool     `gorm:"column:archived;default:false" json:"archived"`
	ExternalID                   *string   `gorm:"column:external_id;size:255" json:"external_id"`
	ToxicityClass                *int      `gorm:"column:toxicity_class" json:"toxicity_class"`
	ActionTerm                   *int      `gorm:"column:action_term" json:"action_term"`
	ActionTermUnits              *string   `gorm:"column:action_term_units;size:255" json:"action_term_units"`
	ActiveSubstance              *string   `gorm:"column:active_substance;size:255" json:"active_substance"`
	DrugForm                     *string   `gorm:"column:drug_form;size:255" json:"drug_form"`
	InfluenceMethod              *string   `gorm:"column:influence_method;size:255" json:"influence_method"`
	BeesIsolatingRecommendedTerm *int      `gorm:"column:bees_isolating_recommended_term" json:"bees_isolating_recommended_term"`
	BeesIsolatingRecommendedTermUnits *string `gorm:"column:bees_isolating_recommended_term_units;size:255" json:"bees_isolating_recommended_term_units"`
	SaleTerm                     *string   `gorm:"column:sale_term;size:255" json:"sale_term"`
	TermOfUse                    *string   `gorm:"column:term_of_use;size:255" json:"term_of_use"`
	WhItemID                     *int64    `gorm:"column:wh_item_id" json:"wh_item_id"`
	WhItemBaseUnitID             *int64    `gorm:"column:wh_item_base_unit_id" json:"wh_item_base_unit_id"`
}

func (Chemical) TableName() string {
	return "chemicals"
}