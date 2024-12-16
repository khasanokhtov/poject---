package models

import "time"

// PersonalIdentifier - модель для таблицы personal_identifiers
type PersonalIdentifier struct {
	ID             int        `json:"id" gorm:"primaryKey"`
	UID            string     `json:"uid"`
	AdditionalInfo string     `json:"additional_info"` // Пустая строка, но не null
	Description    string     `json:"description"`     // Пустая строка, но не null
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	ExternalID     *string    `json:"external_id"`     // Может быть null
}
