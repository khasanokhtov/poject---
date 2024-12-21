package models

import (
	"time"
)

type MachineTask struct {
	ID                            uint       `gorm:"primaryKey;column:id" json:"id"`
	MachineID                     *uint      `gorm:"column:machine_id" json:"machine_id"`
	StartTime                     *time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime                       *time.Time `gorm:"column:end_time" json:"end_time"`
	ActionType                    *string    `gorm:"column:action_type;size:255" json:"action_type"`
	ActionSubtype                 *string    `gorm:"column:action_subtype;size:255" json:"action_subtype"`
	WorkTypeID                    *uint      `gorm:"column:work_type_id" json:"work_type_id"`
	DriverID                      *uint      `gorm:"column:driver_id" json:"driver_id"`
	ImplementID                   *uint      `gorm:"column:implement_id" json:"implement_id"`
	WorkForContractors            *bool      `gorm:"column:work_for_contractors" json:"work_for_contractors"`
	WorkForLandOwners             *bool      `gorm:"column:work_for_land_owners" json:"work_for_land_owners"`
	RealImplementWidth            *float64   `gorm:"column:real_implement_width" json:"real_implement_width"`
	TotalDistance                 *float64   `gorm:"column:total_distance" json:"total_distance"`
	WorkDistance                  *float64   `gorm:"column:work_distance" json:"work_distance"`
	WorkArea                      *float64   `gorm:"column:work_area" json:"work_area"`
	CoveredArea                   *float64   `gorm:"column:covered_area" json:"covered_area"`
	ManualCoveredArea             *float64   `gorm:"column:manual_covered_area" json:"manual_covered_area"`
	WorkDuration                  *float64   `gorm:"column:work_duration" json:"work_duration"`
	AdditionalInfo                *string    `gorm:"column:additional_info" json:"additional_info"`
	Description                   *string    `gorm:"column:description" json:"description"`
	Season                        *int       `gorm:"column:season" json:"season"`
	ExternalID                    *string    `gorm:"column:external_id;size:255" json:"external_id"`
	Status                        *string    `gorm:"column:status;size:255" json:"status"`
	StopsOnRoadDuration           *float64   `gorm:"column:stops_on_road_duration" json:"stops_on_road_duration"`
	MovementsOnRoadDuration       *float64   `gorm:"column:movements_on_road_duration" json:"movements_on_road_duration"`
	TimeWithoutGPSData            *float64   `gorm:"column:time_without_gps_data" json:"time_without_gps_data"`
	FuelConsumption               *float64   `gorm:"column:fuel_consumption" json:"fuel_consumption"`
	FuelConsumptionPerHa          *float64   `gorm:"column:fuel_consumption_per_ha" json:"fuel_consumption_per_ha"`
	FuelConsumptionOnRoad         *float64   `gorm:"column:fuel_consumption_on_road" json:"fuel_consumption_on_road"`
	FuelConsumptionOnRoadAverage  *float64   `gorm:"column:fuel_consumption_on_road_average" json:"fuel_consumption_on_road_average"`
	TrackIntegrityCoef            *float64   `gorm:"column:track_integrity_coef" json:"track_integrity_coef"`
	EngineWorkDurationOnFields    *float64   `gorm:"column:engine_work_duration_on_fields" json:"engine_work_duration_on_fields"`
	EngineWorkDurationOnRoad      *float64   `gorm:"column:engine_work_duration_on_road" json:"engine_work_duration_on_road"`
	LockedToEdit                  *bool      `gorm:"column:locked_to_edit" json:"locked_to_edit"`
	LockedAt                      *time.Time `gorm:"column:locked_at" json:"locked_at"`
	PlanSpeedMin                  *float64   `gorm:"column:plan_speed_min" json:"plan_speed_min"`
	PlanSpeedMax                  *float64   `gorm:"column:plan_speed_max" json:"plan_speed_max"`
	StartFuelLevel                *float64   `gorm:"column:start_fuel_level" json:"start_fuel_level"`
	EndFuelLevel                  *float64   `gorm:"column:end_fuel_level" json:"end_fuel_level"`
	Refuel                        *float64   `gorm:"column:refuel" json:"refuel"`
	CreatedAt                     *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                     *time.Time `gorm:"column:updated_at" json:"updated_at"`
	NotifyResponsibleUsers        *bool      `gorm:"column:notify_responsible_users" json:"notify_responsible_users"`
}

func (MachineTask) TableName() string {
	return "machine_tasks"
}
