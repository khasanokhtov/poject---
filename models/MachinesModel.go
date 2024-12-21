package models

import (
	"time"
)

type Machine struct {
	ID                        int64      `gorm:"primaryKey;column:id" json:"id"`
	Name                      *string    `gorm:"column:name;size:255" json:"name"`
	Model                     *string    `gorm:"column:model;size:255" json:"model"`
	Manufacturer              *string    `gorm:"column:manufacturer;size:255" json:"manufacturer"`
	Year                      *int64     `gorm:"column:year" json:"year"`
	RegistrationNumber        *string    `gorm:"column:registration_number;size:255" json:"registration_number"`
	InventoryNumber           *string    `gorm:"column:inventory_number;size:255" json:"inventory_number"`
	MachineGroupID            *int64     `gorm:"column:machine_group_id" json:"machine_group_id"`
	MachineType               *string    `gorm:"column:machine_type;size:255" json:"machine_type"`
	MachineSubtype            *string    `gorm:"column:machine_subtype;size:255" json:"machine_subtype"`
	AvatarID                  *int64     `gorm:"column:avatar_id" json:"avatar_id"`
	ChassisSerialNumber       *string    `gorm:"column:chassis_serial_number;size:255" json:"chassis_serial_number"`
	EngineSerialNumber        *string    `gorm:"column:engine_serial_number;size:255" json:"engine_serial_number"`
	EnginePower               *float64   `gorm:"column:engine_power" json:"engine_power"`
	FuelType                  *string    `gorm:"column:fuel_type;size:255" json:"fuel_type"`
	FuelTankSize              *float64   `gorm:"column:fuel_tank_size" json:"fuel_tank_size"`
	FuelConsumptionNorm       *float64   `gorm:"column:fuel_consumption_norm" json:"fuel_consumption_norm"`
	LegalCompany              *string    `gorm:"column:legal_company;size:255" json:"legal_company"`
	Description               *string    `gorm:"column:description;type:text" json:"description"`
	DefaultImplementID        *int64     `gorm:"column:default_implement_id" json:"default_implement_id"`
	DefaultDriverID           *int64     `gorm:"column:default_driver_id" json:"default_driver_id"`
	Additional1               *string    `gorm:"column:additional_1;size:255" json:"additional_1"`
	Additional2               *string    `gorm:"column:additional_2;size:255" json:"additional_2"`
	AdditionalInfo            *string    `gorm:"column:additional_info;type:text" json:"additional_info"`
	PhoneNumber               *string    `gorm:"column:phone_number;size:255" json:"phone_number"`
	CreatedAt                 *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 *time.Time `gorm:"column:updated_at" json:"updated_at"`
	ExternalID                *string    `gorm:"column:external_id;size:255" json:"external_id"`
	MachineryModelID          *int64     `gorm:"column:machinery_model_id" json:"machinery_model_id"`
	FuelTypeID                *int64     `gorm:"column:fuel_type_id" json:"fuel_type_id"`
	RefuelSource              *string    `gorm:"column:refuel_source;size:255" json:"refuel_source"`
	MachineryManufacturerID   *int64     `gorm:"column:machinery_manufacturer_id" json:"machinery_manufacturer_id"`
	EngineCapacity            *float64   `gorm:"column:engine_capacity" json:"engine_capacity"`
	Weight                    *float64   `gorm:"column:weight" json:"weight"`
	Height                    *float64   `gorm:"column:height" json:"height"`
	Width                     *float64   `gorm:"column:width" json:"width"`
	Length                    *float64   `gorm:"column:length" json:"length"`
	UnchangedDefaultImplement *bool      `gorm:"column:unchanged_default_implement" json:"unchanged_default_implement"`
	MinDowntimeInSeconds      *int64     `gorm:"column:min_downtime_in_seconds" json:"min_downtime_in_seconds"`
	CalculateDowntimes        *string    `gorm:"column:calculate_downtimes;size:255" json:"calculate_downtimes"`
}

func (Machine) TableName() string {
	return "machines"
}