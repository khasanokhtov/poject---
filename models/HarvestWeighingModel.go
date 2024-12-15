package models

import (
	"time"
)

type HarvestWeighingModel struct {
	ID                      uint       `json:"id" gorm:"primaryKey"`
	MachineID               uint       `json:"machine_id"`
	FieldID                 *uint      `json:"field_id"`
	WeighingPlaceID         uint       `json:"weighing_place_id"`
	TypeOfRoute             *string    `json:"type_of_route"`
	DepartureFromFieldTime  time.Time  `json:"departure_from_field_time"`
	Weight                  float64    `json:"weight"`
	BruttoWeight            float64    `json:"brutto_weight"`
	MarketableWeight        *float64   `json:"marketable_weight"`
	SeedMoisture            *float64   `json:"seed_moisture"`
	SeedAdmixture           *float64   `json:"seed_admixture"`
	WeighingTime            time.Time  `json:"weighing_time"`
	LastTruck               bool       `json:"last_truck"`
	TrackLength             *float64   `json:"track_length"`
	ManuallySetTrackLength  bool       `json:"manually_set_track_length"`
	AdditionalInfo          *string    `json:"additional_info"`
	Description             *string    `json:"description"`
	Season                  uint       `json:"season"`
	CreatedByUserID         uint       `json:"created_by_user_id"`
	ExternalID              *string    `json:"external_id"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
	WaybillNumber           *string    `json:"waybill_number"`
	WaybillDate             *time.Time `json:"waybill_date"`
	HistoryItemID           *uint      `json:"history_item_id"`
	UnloadedMachines        []string   `json:"unloaded_machines" gorm:"-"`
}
