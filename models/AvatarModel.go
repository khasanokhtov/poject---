package models

import (
	"encoding/json"
	"time"
)

type AvatarURL struct {
	URL string `json:"url"`
}

type Avatar struct {
	URL          string    `json:"url"`
	MenuThumb    AvatarURL `json:"menu_thumb"`
	Thumb        AvatarURL `json:"thumb"`
	Small        AvatarURL `json:"small"`
	Tiny         AvatarURL `json:"tiny"`
	SmallRounded AvatarURL `json:"small_rounded"`
	TinyRounded  AvatarURL `json:"tiny_rounded"`
}

type AvatarModel struct {
	ID         int64       `json:"id" gorm:"primaryKey"`
	AvatarType string      `json:"avatar_type"`
	Name       string      `json:"name"`
	Avatar     json.RawMessage `json:"avatar"` // JSON field stored as a raw string
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
