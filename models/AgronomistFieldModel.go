package models

import "time"

type AgronomistFieldModel struct {
    ID        int64     `json:"id" gorm:"primaryKey"`
    UserID    int64     `json:"user_id"`
    FieldID   int64     `json:"field_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
