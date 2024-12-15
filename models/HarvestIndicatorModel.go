package models

import (
	"time"
)

type HarvestIndicatorModel struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	CustomName            *string   `json:"custom_name"`
	MadeAt                time.Time `json:"made_at"`
	SeasonID              uint      `json:"season_id"`
	SeasonName            string    `json:"season_name"`
	HistoryItemID         uint      `json:"history_item_id"`
	HistoryItemName       string    `json:"history_item_name"`
	AgroOperationID       uint      `json:"agro_operation_id"`
	AgroOperationName     string    `json:"agro_operation_name"`
	HarvestWeighingID     *uint     `json:"harvest_weighing_id"`
	HarvestWeighingName   *string   `json:"harvest_weighing_name"`
	GrainHumidity         *float64  `json:"grain_humidity"`
	ProteinContent        *float64  `json:"protein_content"`
	OilContent            *float64  `json:"oil_content"`
	GrainNature           *float64  `json:"grain_nature"`
	HarmfulAdmixture      *float64  `json:"harmful_admixture"`
	GarbageAdmixture      *float64  `json:"garbage_admixture"`
	GrainAdmixture        *float64  `json:"grain_admixture"`
	OilAcidNumber         *float64  `json:"oil_acid_number"`
	GlutenQuality         *float64  `json:"gluten_quality"`
	FallingNumber         *float64  `json:"falling_number"`
	OilByDryMatter        *float64  `json:"oil_by_dry_matter"`
	Core                  *float64  `json:"core"`
	GrainClass            *string   `json:"grain_class"`
	GrainClassName        *string   `json:"grain_class_name"`
	GrainType             *string   `json:"grain_type"`
	GrainTypeName         *string   `json:"grain_type_name"`
	ExternalID            *string   `json:"external_id"`
	AdditionalInfo        *string   `json:"additional_info"`
	Description           *string   `json:"description"`
	Photos                []string  `json:"photos" gorm:"-"`
	PlantThreats          []string  `json:"plant_threats" gorm:"-"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	OilOnDryBasisContent  *float64  `json:"oil_on_dry_basis_content"`
	OilOnWetBasisContent  *float64  `json:"oil_on_wet_basis_content"`
	ErucicAcidContent     *float64  `json:"erucic_acid_content"`
	GlucosinolatesContent *float64  `json:"glucosinolates_content"`
	StarchContent         *float64  `json:"starch_content"`
	GreenIndex            *float64  `json:"green_index"`
}
