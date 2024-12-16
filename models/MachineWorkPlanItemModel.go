package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// MachineWorkPlanItemDate - кастомный тип для обработки даты в формате YYYY-MM-DD
// MachineWorkPlanItemDate - кастомный тип для обработки даты в формате YYYY-MM-DD
type MachineWorkPlanItemDate struct {
	time.Time
}

// UnmarshalJSON - метод для разбора JSON-формата YYYY-MM-DD
func (d *MachineWorkPlanItemDate) UnmarshalJSON(b []byte) error {
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

// MarshalJSON - метод для сериализации в JSON-формат YYYY-MM-DD
func (d MachineWorkPlanItemDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))), nil
}

// Value - метод для преобразования в SQL совместимый тип (реализация Valuer)
func (d MachineWorkPlanItemDate) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}

// Scan - метод для чтения из базы данных (реализация Scanner)
func (d *MachineWorkPlanItemDate) Scan(value interface{}) error {
	if value == nil {
		*d = MachineWorkPlanItemDate{}
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan type %T into MachineWorkPlanItemDate", value)
	}
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

// MachineWorkPlanItemModel - модель для таблицы machine_work_plan_items
type MachineWorkPlanItemModel struct {
    ID                   int                    `json:"id" gorm:"primaryKey"`
    MachineWorkPlanRowID int                    `json:"machine_work_plan_row_id"`
    Date                 MachineWorkPlanItemDate `json:"date"`
    Rate                 float64                `json:"rate"`
    CreatedAt            time.Time              `json:"created_at"`
    UpdatedAt            time.Time              `json:"updated_at"`
}