package models

import (
	"time"
)


type AgroOperation struct {
	ID                          int            `gorm:"primaryKey;column:id"`
	FieldID                     *int           `gorm:"column:field_id"`
	FieldShapeID                *int           `gorm:"column:field_shape_id"`
	AgriWorkPlanID              *int           `gorm:"column:agri_work_plan_id"`
	WorkTypeID                  *int           `gorm:"column:work_type_id"`
	ResponsibleUserIDs          *string        `gorm:"column:responsible_user_ids;type:jsonb"`
	OperationNumber             *string        `gorm:"column:operation_number;size:255"`
	PlannedArea                 *float64       `gorm:"column:planned_area"`
	CompletedArea               *float64       `gorm:"column:completed_area"`
	HarvestedWeight             *float64       `gorm:"column:harvested_weight;default:0"`
	Status                      *string        `gorm:"column:status;size:50"`
	PlannedStartDate            *time.Time     `gorm:"column:planned_start_date"`
	PlannedEndDate              *time.Time     `gorm:"column:planned_end_date"`
	CompletedDate               *time.Time     `gorm:"column:completed_date"`
	ActualStartDatetime         *time.Time     `gorm:"column:actual_start_datetime"`
	CompletedDatetime           *time.Time     `gorm:"column:completed_datetime"`
	Season                      *int           `gorm:"column:season"`
	CustomName                  *string        `gorm:"column:custom_name;size:255"`
	PlannedWaterRate            *float64       `gorm:"column:planned_water_rate"`
	FactWaterRate               *float64       `gorm:"column:fact_water_rate"`
	PlannedRowsSpacing          *float64       `gorm:"column:planned_rows_spacing"`
	PlannedDepth                *float64       `gorm:"column:planned_depth"`
	PlannedSpeed                *float64       `gorm:"column:planned_speed"`
	CompletedPercents           *float64       `gorm:"column:completed_percents"`
	PartiallyCompleted          *bool          `gorm:"column:partially_completed"`
	PartiallyCompletedManuallyDefinedArea *float64 `gorm:"column:partially_completed_manually_defined_area"`
	CoveredArea                 *float64       `gorm:"column:covered_area"`
	CoveredAreaByTrack          *float64       `gorm:"column:covered_area_by_track"`
	MachineWorkArea             *float64       `gorm:"column:machine_work_area"`
	FuelConsumption             *float64       `gorm:"column:fuel_consumption;default:0"`
	FuelConsumptionPerHa        *float64       `gorm:"column:fuel_consumption_per_ha;default:0"`
	AdditionalInfo              *string        `gorm:"column:additional_info;type:text"`
	Description                 *string        `gorm:"column:description;type:text"`
	ApplicationMixItems         *string        `gorm:"column:application_mix_items;type:jsonb"`
	ExternalID                  *string        `gorm:"column:external_id;size:255"`
	CreatedAt                   *time.Time     `gorm:"column:created_at"`
	UpdatedAt                   *time.Time     `gorm:"column:updated_at"`
}

func (AgroOperation) TableName() string {
	return "agro_operations"
}