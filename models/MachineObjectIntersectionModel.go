package models

import (
	"time"
)

// MachineObjectIntersectionModel - структура для хранения данных пересечений объектов машин
type MachineObjectIntersectionModel struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	MachineID             uint      `json:"machine_id"`
	ObjectID              uint      `json:"object_id"`
	ObjectType            string    `json:"object_type"`
	IntersectedObjectType string    `json:"intersected_object_type"`
	IntersectedObjectID   uint      `json:"intersected_object_id"`
	HourStart             time.Time `json:"hour_start"`
	StartIntersection     time.Time `json:"start_intersection"`
	EndIntersection       time.Time `json:"end_intersection"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
