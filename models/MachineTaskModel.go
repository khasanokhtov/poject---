package models

import (
	"time"
)

type MachineTask struct {
	ID                            uint       `gorm:"primaryKey;column:id"`
	MachineID                     *uint      `gorm:"column:machine_id"`
	StartTime                     *time.Time `gorm:"column:start_time"`
	EndTime                       *time.Time `gorm:"column:end_time"`
	ActionType                    *string    `gorm:"column:action_type;size:255"`
	ActionSubtype                 *string    `gorm:"column:action_subtype;size:255"`
	WorkTypeID                    *uint      `gorm:"column:work_type_id"`
	DriverID                      *uint      `gorm:"column:driver_id"`
	ImplementID                   *uint      `gorm:"column:implement_id"`
	WorkForContractors            *bool      `gorm:"column:work_for_contractors"`
	WorkForLandOwners             *bool      `gorm:"column:work_for_land_owners"`
	RealImplementWidth            *float64   `gorm:"column:real_implement_width"`
	TotalDistance                 *float64   `gorm:"column:total_distance"`
	WorkDistance                  *float64   `gorm:"column:work_distance"`
	WorkArea                      *float64   `gorm:"column:work_area"`
	CoveredArea                   *float64   `gorm:"column:covered_area"`
	ManualCoveredArea             *float64   `gorm:"column:manual_covered_area"`
	WorkDuration                  *float64   `gorm:"column:work_duration"`
	AdditionalInfo                *string    `gorm:"column:additional_info"`
	Description                   *string    `gorm:"column:description"`
	Season                        *int       `gorm:"column:season"`
	ExternalID                    *string    `gorm:"column:external_id;size:255"`
	Status                        *string    `gorm:"column:status;size:255"`
	StopsOnRoadDuration           *float64   `gorm:"column:stops_on_road_duration"`
	MovementsOnRoadDuration       *float64   `gorm:"column:movements_on_road_duration"`
	TimeWithoutGPSData            *float64   `gorm:"column:time_without_gps_data"`
	FuelConsumption               *float64   `gorm:"column:fuel_consumption"`
	FuelConsumptionPerHa          *float64   `gorm:"column:fuel_consumption_per_ha"`
	FuelConsumptionOnRoad         *float64   `gorm:"column:fuel_consumption_on_road"`
	FuelConsumptionOnRoadAverage  *float64   `gorm:"column:fuel_consumption_on_road_average"`
	TrackIntegrityCoef            *float64   `gorm:"column:track_integrity_coef"`
	EngineWorkDurationOnFields    *float64   `gorm:"column:engine_work_duration_on_fields"`
	EngineWorkDurationOnRoad      *float64   `gorm:"column:engine_work_duration_on_road"`
	LockedToEdit                  *bool      `gorm:"column:locked_to_edit"`
	LockedAt                      *time.Time `gorm:"column:locked_at"`
	PlanSpeedMin                  *float64   `gorm:"column:plan_speed_min"`
	PlanSpeedMax                  *float64   `gorm:"column:plan_speed_max"`
	StartFuelLevel                *float64   `gorm:"column:start_fuel_level"`
	EndFuelLevel                  *float64   `gorm:"column:end_fuel_level"`
	Refuel                        *float64   `gorm:"column:refuel"`
	CreatedAt                     *time.Time `gorm:"column:created_at"`
	UpdatedAt                     *time.Time `gorm:"column:updated_at"`
	NotifyResponsibleUsers        *bool      `gorm:"column:notify_responsible_users"`
}

func (MachineTask) TableName() string {
	return "machine_tasks"
}
