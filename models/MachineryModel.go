package models

import "time"

// MachineryModel - модель для таблицы machinery_models
type MachineryModel struct {
	ID                    int        `json:"id" gorm:"primaryKey"`
	Name                  string     `json:"name"`
	Hidden                bool       `json:"hidden"`
	ExternalID            *string    `json:"external_id"`            // Может быть null
	MachineryType         string     `json:"machinery_type"`
	AdditionalInfo        *string    `json:"additional_info"`        // Может быть null
	MachineryManufacturerID int      `json:"machinery_manufacturer_id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}
