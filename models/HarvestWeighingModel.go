package models

import (
	"time"
)

type HarvestWeighing struct {
	ID                      int64      `gorm:"primaryKey;column:id" json:"id"`
	MachineID               *int64     `gorm:"column:machine_id" json:"machine_id"`
	FieldID                 *int64     `gorm:"column:field_id" json:"field_id"`
	WeighingPlaceID         *int64     `gorm:"column:weighing_place_id" json:"weighing_place_id"`
	Season                  *int64     `gorm:"column:season" json:"season"`
	DepartureFromFieldTime  *time.Time `gorm:"column:departure_from_field_time" json:"departure_from_field_time"`
	Weight                  *float64   `gorm:"column:weight" json:"weight"`
	BruttoWeigh             *float64   `gorm:"column:brutto_weight" json:"brutto_weight"`
	SeedMoisture            *float64   `gorm:"column:seed_moisture" json:"seed_moisture"`
	SeedAdmixture           *float64   `gorm:"column:seed_admixture" json:"seed_admixture"`
	WeighingTime            *string `gorm:"column:weighing_time" json:"weighing_time"`
	LastTruck               *bool      `gorm:"column:last_truck" json:"last_truck"`
	TrackLength             *float64   `gorm:"column:track_length" json:"track_length"`
	ManuallySetTrackLength  *bool      `gorm:"column:manually_set_track_length" json:"manually_set_track_length"`
	AdditionalInfo          *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	Description             *string    `gorm:"column:description;type:text" json:"description"`
	ExternalID              *string    `gorm:"column:external_id;size:255" json:"external_id"`
	WaybillNumber           *string    `gorm:"column:waybill_number;size:255" json:"waybill_number"`
	WaybillDate             *string `gorm:"column:waybill_date" json:"waybill_date"`
	CreatedByUserID         *int64     `gorm:"column:created_by_user_id" json:"created_by_user_id"`
	UnloadedMachines        []string    `gorm:"column:unloaded_machines;type:jsonb" json:"unloaded_machines"`
	HistoryItemID           *int64     `gorm:"column:history_item_id" json:"history_item_id"`
	MarketableWeight        *float64   `gorm:"column:marketable_weight" json:"marketable_weight"`
	CreatedAt               *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt               *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HarvestWeighing) TableName() string {
	return "harvest_weighings"
}