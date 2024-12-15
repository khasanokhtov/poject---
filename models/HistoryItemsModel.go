package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// CustomDate - кастомный тип для обработки дат в формате "2006-01-02"
type CustomDate struct {
	time.Time
}


// UnmarshalJSON - декодирование JSON для CustomDate
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	// Парсим дату в формате "2006-01-02"
	t, err := time.Parse(customDateFormat, dateStr)
	if err != nil {
		return fmt.Errorf("не удалось разобрать дату: %w", err)
	}

	cd.Time = t
	return nil
}

// MarshalJSON - кодирование CustomDate в JSON
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Time.Format(customDateFormat))
}

// Value - реализация интерфейса driver.Valuer для записи в базу данных
func (cd CustomDate) Value() (driver.Value, error) {
	return cd.Time.Format(customDateFormat), nil
}

// Scan - реализация интерфейса sql.Scanner для чтения из базы данных
func (cd *CustomDate) Scan(value interface{}) error {
	if value == nil {
		cd.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		cd.Time = v
	case string:
		t, err := time.Parse(customDateFormat, v)
		if err != nil {
			return fmt.Errorf("не удалось разобрать дату из строки: %w", err)
		}
		cd.Time = t
	default:
		return fmt.Errorf("неподдерживаемый тип для CustomDate: %T", value)
	}
	return nil
}

// ToTime - конвертация CustomDate в time.Time
func (cd *CustomDate) ToTime() time.Time {
	return cd.Time
}

// HistoryItemModel - структура для хранения данных из API
type HistoryItemModel struct {
	ID                    uint        `json:"id" gorm:"primaryKey"`
	FieldID               uint        `json:"field_id"`
	Year                  uint        `json:"year"`
	Active                bool        `json:"active"`
	CropID                *uint       `json:"crop_id"`
	Variety               *string     `json:"variety"`
	TillType              *string     `json:"till_type"`
	Productivity          *float64    `json:"productivity"`
	ProductivityEstimate  *float64    `json:"productivity_estimate"`
	ProductivityZone      *string     `json:"productivity_zone"`
	SowingDate            *CustomDate `json:"sowing_date"`
	HarvestingDate        *CustomDate `json:"harvesting_date"`
	Description           *string     `json:"description"`
	AdditionalInfo        *string     `json:"additional_info"`
	ExternalID            *string     `json:"external_id"`
	HarvestedWeight       *float64    `json:"harvested_weight"`
	MarketableWeight      *float64    `json:"marketable_weight"`
	YieldDensity          *float64    `json:"yield_density"`
	ExpectedYield         *float64    `json:"expected_yield"`
	AgeOfSugarCane        *float64    `json:"age_of_sugar_cane"`
	GrainClass            *string     `json:"grain_class"`
	GrainHumidity         *float64    `json:"grain_humidity"`
	IrrigationType        *string     `json:"irrigation_type"`
	GrainGarbageAdmixture *float64    `json:"grain_garbage_admixture"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
	ProductionCycleID     *uint       `json:"production_cycle_id"`
	AutomaticEndDate      *CustomDate `json:"automatic_end_date"`
	AutoShapeDetect       bool        `json:"auto_shape_detect"`
	FieldShapeID          *uint       `json:"field_shape_id"`
	SeedID                *uint       `json:"seed_id"`
	PlantsDensity         *float64    `json:"plants_density"`
}
