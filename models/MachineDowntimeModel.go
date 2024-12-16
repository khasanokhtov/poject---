package models

import (
	"time"
)

// MachineDowntimeModel - структура для хранения данных о простоях машин
type MachineDowntimeModel struct {
	ID                   uint       `json:"id" gorm:"primaryKey"`
	StartTime            time.Time  `json:"start_time"`
	EndTime              time.Time  `json:"end_time"`
	DurationInSeconds    int        `json:"duration_in_seconds"`
	MachineDowntimeTypeID *uint     `json:"machine_downtime_type_id"`
	MachineID            uint       `json:"machine_id"`
	Status               string     `json:"status"`
	AdditionalInfo       *string    `json:"additional_info"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	PointLon             float64    `json:"point_lon"`
	PointLat             float64    `json:"point_lat"`
	MachineTaskIDs       []uint     `json:"machine_task_ids" gorm:"-"`
}
