package models

import "time"

// MachineryManufacturerModel - модель для таблицы machinery_manufacturers
type MachineryManufacturerModel struct {
	ID                   int        `json:"id" gorm:"primaryKey"`
	CustomName           *string    `json:"custom_name"`           // Может быть null
	ExternalID           *string    `json:"external_id"`           // Может быть null
	Hidden               bool       `json:"hidden"`
	ImplementManufacturer bool      `json:"implement_manufacturer"`
	MachineManufacturer  bool       `json:"machine_manufacturer"`
	Description          *string    `json:"description"`           // Может быть null
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	StandardName         string     `json:"standard_name"`
}
