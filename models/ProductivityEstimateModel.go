package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// DateOnly - кастомный тип для обработки даты в формате YYYY-MM-DD
type DateOnly struct {
	time.Time
}

// UnmarshalJSON - кастомный метод для разбора формата YYYY-MM-DD
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	str := string(b)
	// Убираем кавычки из строки
	str = str[1 : len(str)-1]
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

// MarshalJSON - метод для сериализации в JSON
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))), nil
}

// Value - метод для преобразования в SQL-совместимый тип
func (d DateOnly) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}

// Scan - метод для чтения из базы данных
func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		*d = DateOnly{}
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan type %T into DateOnly", value)
	}
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

// ProductivityEstimate - модель для таблицы productivity_estimates
type ProductivityEstimate struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	FieldID       int       `json:"field_id"`
	Year          int       `json:"year"`
	EstimateValue float64   `json:"estimate_value"`
	EstimateDate  DateOnly  `json:"estimate_date"` // Используем кастомный тип
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
