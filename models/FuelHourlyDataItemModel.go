package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// FuelDrainData - структура для обработки JSON (объект или массив)
type FuelDrainData struct {
	Data interface{} // Может быть map[string]interface{} или []interface{}
}

// UnmarshalJSON - декодирование JSON в FuelDrainData
func (fd *FuelDrainData) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		fd.Data = nil
		return nil
	}

	// Попробуем декодировать как map[string]interface{}
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err == nil {
		fd.Data = obj
		return nil
	}

	// Если не удалось, попробуем декодировать как массив []interface{}
	var arr []interface{}
	if err := json.Unmarshal(data, &arr); err == nil {
		fd.Data = arr
		return nil
	}

	return fmt.Errorf("ошибка декодирования JSON: %s", string(data))
}

// MarshalJSON - кодирование FuelDrainData в JSON
func (fd FuelDrainData) MarshalJSON() ([]byte, error) {
	return json.Marshal(fd.Data)
}

// Scan - чтение из базы данных
func (fd *FuelDrainData) Scan(value interface{}) error {
	if value == nil {
		fd.Data = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return fd.UnmarshalJSON(v)
	case string:
		return fd.UnmarshalJSON([]byte(v))
	default:
		return fmt.Errorf("неподдерживаемый тип: %T", value)
	}
}

// Value - запись в базу данных
func (fd FuelDrainData) Value() (driver.Value, error) {
	if fd.Data == nil {
		return nil, nil
	}
	return json.Marshal(fd.Data)
}

// FuelHourlyDataItemModel - модель данных
type FuelHourlyDataItemModel struct {
	ID                    uint          `json:"id" gorm:"primaryKey"`
	ObjectType            string        `json:"object_type"`
	ObjectID              uint          `json:"object_id"`
	FuelableType          string        `json:"fuelable_type"`
	FuelableID            uint          `json:"fuelable_id"`
	HourStart             time.Time     `json:"hour_start"`
	FuelConsumption       float64       `json:"fuel_consumption"`
	FuelDrain             FuelDrainData `json:"fuel_drain" gorm:"type:jsonb"` // JSONB в базе
	Refuel                float64       `json:"refuel"`
	DataSourceParameterID uint          `json:"data_source_parameter_id"`
	ExternalID            *string       `json:"external_id"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
}
