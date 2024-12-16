package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// CustomDateMachineRegion - кастомный тип для обработки дат
type CustomDateMachineRegion struct {
	time.Time
}

// UnmarshalJSON - декодирование JSON для CustomDateMachineRegion
func (cd *CustomDateMachineRegion) UnmarshalJSON(data []byte) error {
	var dateStr string

	// Декодируем строку
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	// Если дата пустая, оставляем значение пустым
	if dateStr == "" {
		*cd = CustomDateMachineRegion{}
		return nil
	}

	// Парсим дату в формате "2006-01-02"
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("не удалось разобрать дату: %w", err)
	}

	cd.Time = t
	return nil
}

// MarshalJSON - кодирование CustomDateMachineRegion в JSON
func (cd CustomDateMachineRegion) MarshalJSON() ([]byte, error) {
	if cd.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(cd.Time.Format("2006-01-02"))
}

// Scan - реализация интерфейса sql.Scanner для чтения из базы данных
func (cd *CustomDateMachineRegion) Scan(value interface{}) error {
	if value == nil {
		*cd = CustomDateMachineRegion{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		cd.Time = v
		return nil
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return fmt.Errorf("не удалось разобрать дату: %w", err)
		}
		cd.Time = t
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return fmt.Errorf("не удалось разобрать дату: %w", err)
		}
		cd.Time = t
		return nil
	default:
		return fmt.Errorf("неподдерживаемый тип: %T", value)
	}
}

// Value - реализация интерфейса driver.Valuer для записи в базу данных
func (cd CustomDateMachineRegion) Value() (driver.Value, error) {
	if cd.IsZero() {
		return nil, nil
	}
	return cd.Time.Format("2006-01-02"), nil
}

// MachineRegionMappingItemModel - структура для хранения данных о привязке машин к регионам
type MachineRegionMappingItemModel struct {
	ID             uint                     `json:"id" gorm:"primaryKey"`
	MachineID      uint                     `json:"machine_id"`
	MachineRegionID uint                    `json:"machine_region_id"`
	DateStart      CustomDateMachineRegion  `json:"date_start"`
	DateEnd        *CustomDateMachineRegion `json:"date_end"` // Может быть nil
	NoDateEnd      bool                     `json:"no_date_end"`
	CreatedAt      time.Time                `json:"created_at"`
	UpdatedAt      time.Time                `json:"updated_at"`
}
