package models

import (
	"encoding/json"
	"time"
)

// ProductivityEstimateHistory - модель для таблицы productivity_estimate_histories
type ProductivityEstimateHistory struct {
	ID             int             `json:"id" gorm:"primaryKey"`
	FieldID        int             `json:"field_id"`
	Year           int             `json:"year"`
	EstimateHistory json.RawMessage `json:"estimate_history"` // Сохраняем как JSON
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// GetEstimateHistory - метод для получения EstimateHistory как карты
func (p *ProductivityEstimateHistory) GetEstimateHistory() (map[string]float64, error) {
	var history map[string]float64
	err := json.Unmarshal(p.EstimateHistory, &history)
	return history, err
}

// SetEstimateHistory - метод для установки EstimateHistory из карты
func (p *ProductivityEstimateHistory) SetEstimateHistory(history map[string]float64) error {
	data, err := json.Marshal(history)
	if err != nil {
		return err
	}
	p.EstimateHistory = data
	return nil
}
