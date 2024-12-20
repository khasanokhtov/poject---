package models

import (
	"time"
)

type MachineTaskFieldMappingItem struct {
	ID                          uint        `gorm:"primaryKey;column:id"`
	MachineTaskID               *uint       `gorm:"column:machine_task_id"`
	FieldID                     *uint       `gorm:"column:field_id"`
	CoveredArea                 *float64    `gorm:"column:covered_area"`
	WorkArea                    *float64    `gorm:"column:work_area"`
	CoveredAreaHourly           *string     `gorm:"column:covered_area_hourly;type:jsonb"`
	WorkAreaHourly              *string     `gorm:"column:work_area_hourly;type:jsonb"`
	WorkDistance                *float64    `gorm:"column:work_distance"`
	WorkDistanceHourly          *string     `gorm:"column:work_distance_hourly;type:jsonb"`
	WorkDuration                *float64    `gorm:"column:work_duration"`
	WorkDurationHourly          *string     `gorm:"column:work_duration_hourly;type:jsonb"`
	WorkTimetable               *string     `gorm:"column:work_timetable;type:jsonb"`
	ManuallyDefinedCoveredArea  *bool       `gorm:"column:manually_defined_covered_area"`
	CoveredAreaByTrack          *float64    `gorm:"column:covered_area_by_track"`
	CoveredAreaByTrackHourly    *string     `gorm:"column:covered_area_by_track_hourly;type:jsonb"`
	StopsTimetable              *string     `gorm:"column:stops_timetable;type:jsonb"`
	StopsDuration               *float64    `gorm:"column:stops_duration"`
	StopsDurationHourly         *string     `gorm:"column:stops_duration_hourly;type:jsonb"`
	FuelConsumption             *float64    `gorm:"column:fuel_consumption"`
	FuelConsumptionPerHa        *float64    `gorm:"column:fuel_consumption_per_ha"`
	ManuallyDefinedFuelConsumption *bool    `gorm:"column:manually_defined_fuel_consumption"`
	SensorsData                 *string     `gorm:"column:sensors_data;type:jsonb"`
	InTransit                   *bool       `gorm:"column:in_transit"`
	CreatedAt                   *time.Time  `gorm:"column:created_at"`
	UpdatedAt                   *time.Time  `gorm:"column:updated_at"`
	EngineWorkDuration          *float64    `gorm:"column:engine_work_duration"`
	RawCoveredArea              *float64    `gorm:"column:raw_covered_area"`
	ManualCoveredArea           *float64    `gorm:"column:manual_covered_area"`
	HistoryItemID               *uint       `gorm:"column:history_item_id"`
	FromWarehouseID             *uint       `gorm:"column:from_warehouse_id"`
	LockedToEdit                *bool       `gorm:"column:locked_to_edit"`
	LockedAt                    *time.Time  `gorm:"column:locked_at"`
}

func (MachineTaskFieldMappingItem) TableName() string {
	return "machine_task_field_mapping_items"
}
