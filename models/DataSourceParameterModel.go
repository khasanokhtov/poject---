package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type KeyValue struct {
	Key   float64 `json:"key"`
	Value float64 `json:"value"`
}

type Settings struct {
	FilterMax                     interface{} `json:"filter_max"`
	FilterMin                     interface{} `json:"filter_min"`
	FilterEnabled                 interface{} `json:"filter_enabled"`
	RefuelMinAmount               interface{} `json:"refuel_min_amount"`
	FilterBeforeCalc              interface{} `json:"filter_before_calc"`
	FuelDrainDetectionEnabled     interface{} `json:"fuel_drain_detection_enabled"`
	SensorSmoothingWindowWidth    string `json:"sensor_smoothing_window_width"`
	RefuelSpeedThresholdEnabled   interface{} `json:"refuel_speed_threshold_enabled"`
	SensorSmoothingWindowHeight   string `json:"sensor_smoothing_window_height"`
}

type DataSourceParameterModel struct {
	ID                              uint        `json:"id" gorm:"primaryKey"`
	DataSourceGPSLoggerID           uint        `json:"data_source_gps_logger_id"`
	Name                            string      `json:"name" gorm:"not null"`
	NameHuman                       string      `json:"name_human"`
	UnitsOfMeasurement              string      `json:"units_of_measurement"`
	ValueType                       string      `json:"value_type"`
	CalcType                        string      `json:"calc_type"`
	GPSSensorName                   string      `json:"gps_sensor_name"`
	CalcFormula                     *string     `json:"calc_formula"`
	KeyValuesTable                  [][]float64 `json:"key_values_table" gorm:"-"`
	KeyValuesTableJSON              string      `json:"-" gorm:"column:key_values_table"`
	Description                     string      `json:"description"`
	Hidden                          bool        `json:"hidden"`
	PresentationMode                string      `json:"presentation_mode"`
	Type                            string      `json:"type"`
	AutomaticallyAssignInMachineTasks bool      `json:"automatically_assign_in_machine_tasks"`
	FuelConsumptionAccounting       bool        `json:"fuel_consumption_accounting"`
	FuelFlowSensorType              string      `json:"fuel_flow_sensor_type"`
	Settings                        Settings    `json:"settings" gorm:"-"`
	CreatedAt                       time.Time   `json:"created_at"`
	UpdatedAt                       time.Time   `json:"updated_at"`
}

func (d *DataSourceParameterModel) BeforeSave(tx *gorm.DB) (err error) {
	if len(d.KeyValuesTable) > 0 {
		// Сериализуем массив массивов в JSON
		jsonData, err := json.Marshal(d.KeyValuesTable)
		if err != nil {
			return fmt.Errorf("ошибка сериализации KeyValuesTable: %w", err)
		}
		d.KeyValuesTableJSON = string(jsonData) // Устанавливаем сериализованные данные
	}
	return nil
}