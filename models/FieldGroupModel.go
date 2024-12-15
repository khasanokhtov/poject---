package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type customDate time.Time

const customDateFormat = "2006-01-02" // Формат "YYYY-MM-DD"

// UnmarshalJSON - для декодирования JSON
func (cd *customDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) > 1 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	parsedTime, err := time.Parse(customDateFormat, s)
	if err != nil {
		return fmt.Errorf("ошибка парсинга даты: %w", err)
	}
	*cd = customDate(parsedTime)
	return nil
}

// MarshalJSON - для кодирования JSON
func (cd customDate) MarshalJSON() ([]byte, error) {
	t := time.Time(cd)
	return []byte(fmt.Sprintf(`"%s"`, t.Format(customDateFormat))), nil
}

// ToTime - преобразование в time.Time
func (cd customDate) ToTime() time.Time {
	return time.Time(cd)
}

// Scan - реализация интерфейса sql.Scanner для чтения из базы данных
func (cd *customDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*cd = customDate(v)
		return nil
	case []byte:
		parsedTime, err := time.Parse(customDateFormat, string(v))
		if err != nil {
			return fmt.Errorf("ошибка парсинга даты: %w", err)
		}
		*cd = customDate(parsedTime)
		return nil
	case string:
		parsedTime, err := time.Parse(customDateFormat, v)
		if err != nil {
			return fmt.Errorf("ошибка парсинга даты: %w", err)
		}
		*cd = customDate(parsedTime)
		return nil
	}
	return fmt.Errorf("неподдерживаемый тип: %T", value)
}

// Value - реализация интерфейса driver.Valuer для записи в базу данных
func (cd customDate) Value() (driver.Value, error) {
	t := time.Time(cd)
	return t.Format(customDateFormat), nil
}



type FieldGroupModel struct {
	ID                          uint      `json:"id" gorm:"primaryKey"`
	GroupFolderID               uint      `json:"group_folder_id"`
	Name                        string    `json:"name" gorm:"not null"`
	Description                 string    `json:"description"`
	AdministrativeAreaName      string    `json:"administrative_area_name"`
	SubAdministrativeAreaName   string    `json:"subadministrative_area_name"`
	Locality                    string    `json:"locality"`
	Hidden                      bool      `json:"hidden"`
	ExternalID                  *string   `json:"external_id"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	IdempotencyKey              *string   `json:"idempotency_key"`
	LegalEntity                 string    `json:"legal_entity"`
	MachineTaskDefaultDuration  uint      `json:"machine_task_default_duration"`
	AccountingPeriodClosingDate customDate `json:"accounting_period_closing_date" gorm:"type:date"`
	MachineTaskDefaultStartTime string    `json:"machine_task_default_start_time"`
}

