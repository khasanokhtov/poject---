package models

import "time"

type FertilizerElements struct {
	N   float64 `json:"N"`
	P2O5 float64 `json:"P2O5"`
	K2O float64 `json:"K2O"`
	Ca  float64 `json:"Ca"`
	Mg  float64 `json:"Mg"`
	S   float64 `json:"S"`
	B   float64 `json:"B"`
	Cl  float64 `json:"Cl"`
	Cu  float64 `json:"Cu"`
	Fe  float64 `json:"Fe"`
	Mn  float64 `json:"Mn"`
	Mo  float64 `json:"Mo"`
	Ni  float64 `json:"Ni"`
	Zn  float64 `json:"Zn"`
	Co  float64 `json:"Co"`
	Se  float64 `json:"Se"`
}

type FertilizerModel struct {
	ID                    uint            `json:"id" gorm:"primaryKey"`
	Name                  string          `json:"name" gorm:"not null"`
	ManufacturerName      *string         `json:"manufacturer_name"`
	FertilizerType        string          `json:"fertilizer_type"`
	SourceType            string          `json:"source_type"`
	NutrientType          string          `json:"nutrient_type"`
	UnitsOfMeasurement    string          `json:"units_of_measurement"`
	Elements              FertilizerElements `json:"elements" gorm:"embedded;embeddedPrefix:element_"`
	Archived              bool            `json:"archived"`
	WhItemID              *uint           `json:"wh_item_id"`
	WhItemBaseUnitID      *uint           `json:"wh_item_base_unit_id"`
	BaseInventoryUnitID   *uint           `json:"base_inventory_unit_id"`
	AdditionalInfo        *string         `json:"additional_info"`
	Description           *string         `json:"description"`
	Density               *string         `json:"density"`
	ExternalID            *string         `json:"external_id"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
}
