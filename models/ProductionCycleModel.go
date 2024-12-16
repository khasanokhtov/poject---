package models

import "time"

// ProductionCycle - модель для таблицы production_cycles
type ProductionCycle struct {
	ID             int        `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	CropID         *int       `json:"crop_id"`          // Может быть null
	SeasonID       int        `json:"season_id"`
	BaseCycle      bool       `json:"base_cycle"`
	AdditionalInfo *string    `json:"additional_info"`  // Может быть null
	Description    *string    `json:"description"`      // Может быть null
	ExternalID     *string    `json:"external_id"`      // Может быть null
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
