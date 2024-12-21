package models

import (
	"encoding/json"
	"time"
)

type MachineTaskFieldMappingItem struct {
	ID                          uint        		`gorm:"primaryKey;column:id" json:"id"`
	MachineTaskID               *uint       		`gorm:"column:machine_task_id" json:"machine_task_id"`
	FieldID                     *uint       		`gorm:"column:field_id" json:"field_id"`
	CoveredArea                 *float64    		`gorm:"column:covered_area" json:"covered_area"`
	WorkArea                    *float64    		`gorm:"column:work_area" json:"work_area"`
	CoveredAreaHourly           json.RawMessage     `gorm:"column:covered_area_hourly;type:jsonb" json:"covered_area_hourly"`
	WorkAreaHourly              json.RawMessage     `gorm:"column:work_area_hourly;type:jsonb" json:"work_area_hourly"`
	WorkDistance                *float64    		`gorm:"column:work_distance" json:"work_distance"`
	WorkDistanceHourly          json.RawMessage     `gorm:"column:work_distance_hourly;type:jsonb" json:"work_distance_hourly"`
	WorkDuration                *float64    		`gorm:"column:work_duration" json:"work_duration"`
	WorkDurationHourly          json.RawMessage     `gorm:"column:work_duration_hourly;type:jsonb" json:"work_duration_hourly"`
	WorkTimetable               json.RawMessage     `gorm:"column:work_timetable;type:jsonb" json:"work_timetable"`
	ManuallyDefinedCoveredArea  *bool       		`gorm:"column:manually_defined_covered_area" json:"manually_defined_covered_area"`
	CoveredAreaByTrack          *float64    		`gorm:"column:covered_area_by_track" json:"covered_area_by_track"`
	CoveredAreaByTrackHourly    json.RawMessage     `gorm:"column:covered_area_by_track_hourly;type:jsonb" json:"covered_area_by_track_hourly"`
	StopsTimetable              json.RawMessage     `gorm:"column:stops_timetable;type:jsonb" json:"stops_timetable"`
	StopsDuration               *float64    		`gorm:"column:stops_duration" json:"stops_duration"`
	StopsDurationHourly         json.RawMessage     `gorm:"column:stops_duration_hourly;type:jsonb" json:"stops_duration_hourly"`
	FuelConsumption             *float64    		`gorm:"column:fuel_consumption" json:"fuel_consumption"`
	FuelConsumptionPerHa        *float64    		`gorm:"column:fuel_consumption_per_ha" json:"fuel_consumption_per_ha"`
	ManuallyDefinedFuelConsumption *bool    		`gorm:"column:manually_defined_fuel_consumption" json:"manually_defined_fuel_consumption"`
	SensorsData                 json.RawMessage     `gorm:"column:sensors_data;type:jsonb" json:"sensors_data"`
	InTransit                   *bool       		`gorm:"column:in_transit" json:"in_transit"`
	CreatedAt                   *time.Time  		`gorm:"column:created_at" json:"created_at"`
	UpdatedAt                   *time.Time  		`gorm:"column:updated_at" json:"updated_at"`
	EngineWorkDuration          *float64    		`gorm:"column:engine_work_duration" json:"engine_work_duration"`
	RawCoveredArea              *float64    		`gorm:"column:raw_covered_area" json:"raw_covered_area"`
	ManualCoveredArea           *float64    		`gorm:"column:manual_covered_area" json:"manual_covered_area"`
	HistoryItemID               *uint       		`gorm:"column:history_item_id" json:"history_item_id"`
	FromWarehouseID             *uint       		`gorm:"column:from_warehouse_id" json:"from_warehouse_id"`
	LockedToEdit                *bool       		`gorm:"column:locked_to_edit" json:"locked_to_edit"`
	LockedAt                    *time.Time  		`gorm:"column:locked_at" json:"locked_at"`
}

func (MachineTaskFieldMappingItem) TableName() string {
	return "machine_task_field_mapping_items"
}
