package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// CustomDateInventory - кастомный тип для обработки дат с пустыми значениями
type CustomDateInventory struct {
	time.Time
}

// UnmarshalJSON - декодирование JSON для CustomDateInventory с обработкой пустых значений
func (cd *CustomDateInventory) UnmarshalJSON(data []byte) error {
	var dateStr string

	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	// Если дата пустая или null, оставляем значение пустым
	if dateStr == "" || dateStr == "null" {
		*cd = CustomDateInventory{}
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

// MarshalJSON - кодирование CustomDateInventory в JSON
func (cd CustomDateInventory) MarshalJSON() ([]byte, error) {
	if cd.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(cd.Time.Format("2006-01-02"))
}

// Value - реализация интерфейса driver.Valuer для записи в базу данных
func (cd CustomDateInventory) Value() (driver.Value, error) {
	if cd.IsZero() {
		return nil, nil
	}
	return cd.Time.Format("2006-01-02"), nil
}

// Scan - реализация интерфейса sql.Scanner для чтения из базы данных
func (cd *CustomDateInventory) Scan(value interface{}) error {
	if value == nil {
		*cd = CustomDateInventory{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		cd.Time = v
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return fmt.Errorf("не удалось разобрать дату из строки: %w", err)
		}
		cd.Time = t
	default:
		return fmt.Errorf("неподдерживаемый тип для даты: %T", value)
	}
	return nil
}

// InventoryHistoryItemModel - структура для хранения истории инвентаризации
type InventoryHistoryItemModel struct {
	ID              uint               `json:"id" gorm:"primaryKey"`
	HistoryableID   uint               `json:"historyable_id"`
	HistoryableType string             `json:"historyable_type"`
	EventStartAt    CustomDateInventory `json:"event_start_at"`
	EventEndAt      CustomDateInventory `json:"event_end_at"`
	Reason          string             `json:"reason"`
	Description     *string            `json:"description"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	Available       bool               `json:"available"`
	Hidden          bool               `json:"hidden"`
	ExternalID      *string            `json:"external_id"`
}
