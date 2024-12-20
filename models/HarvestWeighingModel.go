package models

import (
	"time"
)

type HarvestWeighing struct {
	ID                      int64      `gorm:"primaryKey;column:id"`
	MachineID               *int64     `gorm:"column:machine_id"`
	FieldID                 *int64     `gorm:"column:field_id"`
	WeighingPlaceID         *int64     `gorm:"column:weighing_place_id"`
	Season                  *int64     `gorm:"column:season"`
	DepartureFromFieldTime  *time.Time `gorm:"column:departure_from_field_time"`
	Weight                  *float64   `gorm:"column:weight"`
	BruttoWeigh             *float64   `gorm:"column:brutto_weigh"`
	SeedMoisture            *float64   `gorm:"column:seed_moisture"`
	SeedAdmixture           *float64   `gorm:"column:seed_admixture"`
	WeighingTime            *time.Time `gorm:"column:weighing_time"`
	LastTruck               *bool      `gorm:"column:last_truck"`
	TrackLength             *float64   `gorm:"column:track_length"`
	ManuallySetTrackLength  *bool      `gorm:"column:manually_set_track_length"`
	AdditionalInfo          *string    `gorm:"column:additional_info;type:text"`
	Description             *string    `gorm:"column:description;type:text"`
	ExternalID              *string    `gorm:"column:external_id;size:255"`
	WaybillNumber           *string    `gorm:"column:waybill_number;size:255"`
	WaybillDate             *time.Time `gorm:"column:waybill_date"`
	CreatedByUserID         *int64     `gorm:"column:created_by_user_id"`
	UnloadedMachines        *string    `gorm:"column:unloaded_machines;type:jsonb"`
	HistoryItemID           *int64     `gorm:"column:history_item_id"`
	MarketableWeight        *float64   `gorm:"column:marketable_weight"`
	CreatedAt               *time.Time `gorm:"column:created_at"`
	UpdatedAt               *time.Time `gorm:"column:updated_at"`
}

func (HarvestWeighing) TableName() string {
	return "harvest_weighings"
}