package models

import "time"

// OdometerState - модель для таблицы odometer_states
type OdometerState struct {
	ID                        int        `json:"id" gorm:"primaryKey"`
	MachineID                 int        `json:"machine_id"`
	OdometerableID            int        `json:"odometerable_id"`
	OdometerableType          string     `json:"odometerable_type"`
	ManualDistance            *float64   `json:"manual_distance"`              // Может быть null
	ManualEngineWorkDuration  *float64   `json:"manual_engine_work_duration"`  // Может быть null
	CalculatedDistance        float64    `json:"calculated_distance"`
	CalculatedEngineWorkDuration *float64 `json:"calculated_engine_work_duration"` // Может быть null
	MeasuringTime             time.Time  `json:"measuring_time"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
}
