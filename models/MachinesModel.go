package models

import (
	"time"
)

type Machine struct {
	ID                        int64      `gorm:"primaryKey;column:id"`
	Name                      *string    `gorm:"column:name;size:255"`
	Model                     *string    `gorm:"column:model;size:255"`
	Manufacturer              *string    `gorm:"column:manufacturer;size:255"`
	Year                      *int64     `gorm:"column:year"`
	RegistrationNumber        *string    `gorm:"column:registration_number;size:255"`
	InventoryNumber           *string    `gorm:"column:inventory_number;size:255"`
	MachineGroupID            *int64     `gorm:"column:machine_group_id"`
	MachineType               *string    `gorm:"column:machine_type;size:255"`
	MachineSubtype            *string    `gorm:"column:machine_subtype;size:255"`
	AvatarID                  *int64     `gorm:"column:avatar_id"`
	ChassisSerialNumber       *string    `gorm:"column:chassis_serial_number;size:255"`
	EngineSerialNumber        *string    `gorm:"column:engine_serial_number;size:255"`
	EnginePower               *float64   `gorm:"column:engine_power"`
	FuelType                  *string    `gorm:"column:fuel_type;size:255"`
	FuelTankSize              *float64   `gorm:"column:fuel_tank_size"`
	FuelConsumptionNorm       *float64   `gorm:"column:fuel_consumption_norm"`
	LegalCompany              *string    `gorm:"column:legal_company;size:255"`
	Description               *string    `gorm:"column:description;type:text"`
	DefaultImplementID        *int64     `gorm:"column:default_implement_id"`
	DefaultDriverID           *int64     `gorm:"column:default_driver_id"`
	Additional1               *string    `gorm:"column:additional_1;size:255"`
	Additional2               *string    `gorm:"column:additional_2;size:255"`
	AdditionalInfo            *string    `gorm:"column:additional_info;type:text"`
	PhoneNumber               *string    `gorm:"column:phone_number;size:255"`
	CreatedAt                 *time.Time `gorm:"column:created_at"`
	UpdatedAt                 *time.Time `gorm:"column:updated_at"`
	ExternalID                *int64     `gorm:"column:external_id"`
	MachineryModelID          *int64     `gorm:"column:machinery_model_id"`
	FuelTypeID                *int64     `gorm:"column:fuel_type_id"`
	RefuelSource              *string    `gorm:"column:refuel_source;size:255"`
	MachineryManufacturerID   *int64     `gorm:"column:machinery_manufacturer_id"`
	EngineCapacity            *float64   `gorm:"column:engine_capacity"`
	Weight                    *float64   `gorm:"column:weight"`
	Height                    *float64   `gorm:"column:height"`
	Width                     *float64   `gorm:"column:width"`
	Length                    *float64   `gorm:"column:length"`
	UnchangedDefaultImplement *bool      `gorm:"column:unchanged_default_implement"`
	MinDowntimeInSeconds      *int64     `gorm:"column:min_downtime_in_seconds"`
	CalculateDowntimes        *string    `gorm:"column:calculate_downtimes;size:255"`
}

func (Machine) TableName() string {
	return "machines"
}