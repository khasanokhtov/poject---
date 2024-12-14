package models

import (
	"time"
)

// Модель для хранения активных веществ
type ActiveSubstance struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Основная модель Chemical
type ChemicalModel struct {
	ID                           int64           `json:"id" gorm:"primaryKey"`
	Name                         string          `json:"name"`
	ManufacturerName             *string         `json:"manufacturer_name"`
	ChemicalType                 string          `json:"chemical_type"`
	UnitsOfMeasurement           string          `json:"units_of_measurement"`
	AdditionalInfo               *string         `json:"additional_info"`
	Description                  *string         `json:"description"`
	CreatedAt                    time.Time       `json:"created_at"`
	UpdatedAt                    time.Time       `json:"updated_at"`
	Archived                     bool            `json:"archived"`
	ExternalID                   *string         `json:"external_id"`
	ToxicityClass                *int            `json:"toxicity_class"`
	ActionTerm                   *string         `json:"action_term"`
	ActionTermUnits              *string         `json:"action_term_units"`
	ActiveSubstance              *string         `json:"active_substance"`
	DrugForm                     *string         `json:"drug_form"`
	InfluenceMethod              *string         `json:"influence_method"`
	BeesIsolatingRecommendedTerm *int            `json:"bees_isolating_recommended_term"`
	BeesIsolatingRecommendedTermUnits *string    `json:"bees_isolating_recommended_term_units"`
	SaleTerm                     *string         `json:"sale_term"`
	TermOfUse                    *string         `json:"term_of_use"`
	WhItemID                     int64           `json:"wh_item_id"`
	WhItemBaseUnitID             int64           `json:"wh_item_base_unit_id"`
	BaseInventoryUnitID          int64           `json:"base_inventory_unit_id"`
	ActiveSubstances             []ActiveSubstance `json:"active_substances" gorm:"-"`
}
