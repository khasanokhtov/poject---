package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// FuelDrainData - структура для обработки JSON объекта
type FuelDrainData map[string]interface{}

// UnmarshalJSON - декодирование из JSON
func (fd *FuelDrainData) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        *fd = nil
        return nil
    }
    var obj map[string]interface{}
    if err := json.Unmarshal(data, &obj); err != nil {
        return fmt.Errorf("ошибка декодирования JSON: %w", err)
    }
    *fd = obj
    return nil
}

// MarshalJSON - кодирование в JSON
func (fd FuelDrainData) MarshalJSON() ([]byte, error) {
    if fd == nil {
        return json.Marshal(nil)
    }
    return json.Marshal(map[string]interface{}(fd))
}

// Scan - чтение из базы данных
func (fd *FuelDrainData) Scan(value interface{}) error {
    if value == nil {
        *fd = nil
        return nil
    }

    switch v := value.(type) {
    case []byte:
        var obj map[string]interface{}
        if err := json.Unmarshal(v, &obj); err != nil {
            return fmt.Errorf("ошибка декодирования JSON: %w", err)
        }
        *fd = obj
        return nil
    case string:
        var obj map[string]interface{}
        if err := json.Unmarshal([]byte(v), &obj); err != nil {
            return fmt.Errorf("ошибка декодирования строки JSON: %w", err)
        }
        *fd = obj
        return nil
    default:
        return fmt.Errorf("неподдерживаемый тип: %T", value)
    }
}

// Value - запись в базу данных
func (fd FuelDrainData) Value() (driver.Value, error) {
    if fd == nil {
        return nil, nil
    }
    return json.Marshal(fd)
}

// FuelHourlyDataItemModel - модель данных
type FuelHourlyDataItemModel struct {
    ID                   uint          `json:"id" gorm:"primaryKey"`
    ObjectType           string        `json:"object_type"`
    ObjectID             uint          `json:"object_id"`
    FuelableType         string        `json:"fuelable_type"`
    FuelableID           uint          `json:"fuelable_id"`
    HourStart            time.Time     `json:"hour_start"`
    FuelConsumption      float64       `json:"fuel_consumption"`
    FuelDrain            FuelDrainData `json:"fuel_drain" gorm:"type:jsonb"` // JSONB в базе
    Refuel               float64       `json:"refuel"`
    DataSourceParameterID uint         `json:"data_source_parameter_id"`
    ExternalID           *string       `json:"external_id"`
    CreatedAt            time.Time     `json:"created_at"`
    UpdatedAt            time.Time     `json:"updated_at"`
}
