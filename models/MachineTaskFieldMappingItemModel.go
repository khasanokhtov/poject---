package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// HourlyData - универсальная структура для хранения данных почасовых массивов
type HourlyData []struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
}

// UnmarshalJSON - декодирование JSON массива в HourlyData
func (hd *HourlyData) UnmarshalJSON(data []byte) error {
	var raw [][]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var parsed HourlyData
	for _, entry := range raw {
		if len(entry) == 2 {
			if timestamp, ok := entry[0].(string); ok {
				if value, ok := entry[1].(float64); ok {
					parsed = append(parsed, struct {
						Timestamp string  `json:"timestamp"`
						Value     float64 `json:"value"`
					}{
						Timestamp: timestamp,
						Value:     value,
					})
				}
			}
		}
	}
	*hd = parsed
	return nil
}

// Value - запись в базу данных
func (hd HourlyData) Value() (driver.Value, error) {
	return json.Marshal(hd)
}

// WorkTimetableData - структура для расписаний
type WorkTimetableData [][]string

// Value - запись в базу данных
func (wtd WorkTimetableData) Value() (driver.Value, error) {
	return json.Marshal(wtd)
}

// Scan - чтение из базы данных
func (wtd *WorkTimetableData) Scan(value interface{}) error {
	if value == nil {
		*wtd = nil
		return nil
	}
	return json.Unmarshal(value.([]byte), wtd)
}

// MachineTaskFieldMappingItemModel - основная модель
type MachineTaskFieldMappingItemModel struct {
	ID                        uint               `json:"id" gorm:"primaryKey"`
	MachineTaskID             uint               `json:"machine_task_id"`
	FieldID                   uint               `json:"field_id"`
	CoveredArea               float64            `json:"covered_area"`
	WorkArea                  float64            `json:"work_area"`
	CoveredAreaHourly         HourlyData         `json:"covered_area_hourly" gorm:"type:jsonb"`
	WorkAreaHourly            HourlyData         `json:"work_area_hourly" gorm:"type:jsonb"`
	WorkDistance              float64            `json:"work_distance"`
	WorkDistanceHourly        HourlyData         `json:"work_distance_hourly" gorm:"type:jsonb"`
	WorkDuration              uint               `json:"work_duration"`
	WorkDurationHourly        HourlyData         `json:"work_duration_hourly" gorm:"type:jsonb"`
	WorkTimetable             WorkTimetableData  `json:"work_timetable" gorm:"type:jsonb"`
	ManuallyDefinedCoveredArea bool              `json:"manually_defined_covered_area"`
	CoveredAreaByTrack        float64            `json:"covered_area_by_track"`
	CoveredAreaByTrackHourly  HourlyData         `json:"covered_area_by_track_hourly" gorm:"type:jsonb"`
	StopsTimetable            WorkTimetableData  `json:"stops_timetable" gorm:"type:jsonb"`
	StopsDuration             uint               `json:"stops_duration"`
	StopsDurationHourly       HourlyData         `json:"stops_duration_hourly" gorm:"type:jsonb"`
	FuelConsumption           float64            `json:"fuel_consumption"`
	FuelConsumptionPerHa      float64            `json:"fuel_consumption_per_ha"`
	ManuallyDefinedFuelConsumption bool          `json:"manually_defined_fuel_consumption"`
	InTransit                 bool               `json:"in_transit"`
	CreatedAt                 time.Time          `json:"created_at"`
	UpdatedAt                 time.Time          `json:"updated_at"`
	EngineWorkDuration        uint               `json:"engine_work_duration"`
	RawCoveredArea            float64            `json:"raw_covered_area"`
	ManualCoveredArea         *float64           `json:"manual_covered_area"`
	HistoryItemID             uint               `json:"history_item_id"`
	FromWarehouseID           *uint              `json:"from_warehouse_id"`
	LockedToEdit              *bool              `json:"locked_to_edit"`
	LockedAt                  *time.Time         `json:"locked_at"`
}
