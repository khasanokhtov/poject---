package models

import (
	"time"
)

// MachineTaskModel - обновлённая модель данных для machine_tasks
type MachineTaskModel struct {
	ID                            uint       `json:"id" gorm:"primaryKey"`
	MachineID                     uint       `json:"machine_id"`
	StartTime                     time.Time  `json:"start_time"`
	EndTime                       time.Time  `json:"end_time"`
	ActionType                    string     `json:"action_type"`
	ActionSubtype                 string     `json:"action_subtype"`
	WorkTypeID                    uint       `json:"work_type_id"`
	DriverID                      uint       `json:"driver_id"`
	ImplementID                   uint       `json:"implement_id"`
	WorkForContractors            bool       `json:"work_for_contractors"`
	WorkForLandOwners             bool       `json:"work_for_land_owners"`
	RealImplementWidth            *float64   `json:"real_implement_width"`
	TotalDistance                 float64    `json:"total_distance"`
	WorkDistance                  float64    `json:"work_distance"`
	WorkArea                      float64    `json:"work_area"`
	CoveredArea                   float64    `json:"covered_area"`
	ManualCoveredArea             *float64   `json:"manual_covered_area"`
	WorkDuration                  int        `json:"work_duration"`
	AdditionalInfo                string     `json:"additional_info"`
	Description                   string     `json:"description"`
	Season                        int        `json:"season"`
	ExternalID                    *string    `json:"external_id"`
	Status                        string     `json:"status"`
	StopsOnRoadDuration           int        `json:"stops_on_road_duration"`
	MovementsOnRoadDuration       int        `json:"movements_on_road_duration"`
	TimeWithoutGPSData            int        `json:"time_without_gps_data"`
	FuelConsumption               float64    `json:"fuel_consumption"`
	FuelConsumptionPerHa          float64    `json:"fuel_consumption_per_ha"`
	FuelConsumptionOnRoad         float64    `json:"fuel_consumption_on_road"`
	FuelConsumptionOnRoadAverage  float64    `json:"fuel_consumption_on_road_average"`
	TrackIntegrityCoef            float64    `json:"track_integrity_coef"`
	EngineWorkDurationOnFields    int        `json:"engine_work_duration_on_fields"`
	EngineWorkDurationOnRoad      int        `json:"engine_work_duration_on_road"`
	LockedToEdit                  *bool      `json:"locked_to_edit"`
	LockedAt                      *time.Time `json:"locked_at"`
	PlanSpeedMin                  *float64   `json:"plan_speed_min"`
	PlanSpeedMax                  *float64   `json:"plan_speed_max"`
	StartFuelLevel                *float64   `json:"start_fuel_level"`
	EndFuelLevel                  *float64   `json:"end_fuel_level"`
	Refuel                        *float64   `json:"refuel"`
	CreatedAt                     time.Time  `json:"created_at"`
	UpdatedAt                     time.Time  `json:"updated_at"`
	NotifyResponsibleUsers        bool       `json:"notify_responsible_users"`
}
