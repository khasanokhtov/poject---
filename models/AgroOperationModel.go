package models

import (
	"time"
)


type AgroOperations struct {
	ID                          int        `gorm:"primaryKey;column:id" json:"id"`
	FieldID                     *int       `gorm:"column:field_id" json:"field_id"`
	FieldShapeID                *int       `gorm:"column:field_shape_id" json:"field_shape_id"`
	AgriWorkPlanID              *int       `gorm:"column:agri_work_plan_id" json:"agri_work_plan_id"`
	WorkTypeID                  *int       `gorm:"column:work_type_id" json:"work_type_id"`
	OperationNumber             *string    `gorm:"column:operation_number;size:255" json:"operation_number"`
	PlannedArea                 *float64   `gorm:"column:planned_area" json:"planned_area"`
	CompletedArea               *float64   `gorm:"column:completed_area" json:"completed_area"`
	HarvestedWeight             *float64   `gorm:"column:harvested_weight;default:0" json:"harvested_weight"`
	Status                      *string    `gorm:"column:status;size:50" json:"status"`
	PlannedStartDate            *string    `gorm:"column:planned_start_date" json:"planned_start_date"`
	PlannedEndDate              *string    `gorm:"column:planned_end_date" json:"planned_end_date"`
	CompletedDate               *string    `gorm:"column:completed_date" json:"completed_date"`
	ActualStartDatetime         *string    `gorm:"column:actual_start_datetime" json:"actual_start_datetime"`
	CompletedDatetime           *string    `gorm:"column:completed_datetime" json:"completed_datetime"`
	Season                      *int       `gorm:"column:season" json:"season"`
	CustomName                  *string    `gorm:"column:custom_name;size:255" json:"custom_name"`
	PlannedWaterRate            *float64   `gorm:"column:planned_water_rate" json:"planned_water_rate"`
	FactWaterRate               *float64   `gorm:"column:fact_water_rate" json:"fact_water_rate"`
	PlannedRowsSpacing          *float64   `gorm:"column:planned_rows_spacing" json:"planned_row_spacing"`
	PlannedDepth                *float64    `gorm:"column:planned_depth" json:"planned_depth"`
	PlannedSpeed                *float64    `gorm:"column:planned_speed" json:"planned_speed"`
	CompletedPercents           *float64   `gorm:"column:completed_percents" json:"completed_percents"`
	PartiallyCompleted          *bool      `gorm:"column:partially_completed" json:"partially_completed"`
	PartiallyCompletedManuallyDefinedArea *float64 `gorm:"column:partially_completed_manually_defined_area" json:"partially_completed_manually_defined_area"`
	CoveredArea                 *float64   `gorm:"column:covered_area" json:"covered_area"`
	CoveredAreaByTrack          *float64   `gorm:"column:covered_area_by_track" json:"covered_area_by_track"`
	MachineWorkArea             *float64   `gorm:"column:machine_work_area" json:"machine_work_area"`
	FuelConsumption             *float64   `gorm:"column:fuel_consumption;default:0" json:"fuel_consumption"`
	FuelConsumptionPerHa        *float64   `gorm:"column:fuel_consumption_per_ha;default:0" json:"fuel_consumption_per_ha"`
	AdditionalInfo              *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	Description                 *string    `gorm:"column:description;type:text" json:"description"`
	ApplicationMixItems         *string    `gorm:"column:application_mix_items;type:jsonb" json:"application_mix_items"`
	ExternalID                  *string    `gorm:"column:external_id;size:255" json:"external_id"`
	CreatedAt                   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (AgroOperations) TableName() string {
	return "agro_operations"
}