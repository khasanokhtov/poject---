package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

const customDateFormat = "2006-01-02"

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	dateStr := string(data)
	dateStr = dateStr[1 : len(dateStr)-1] // Убираем кавычки
	parsedTime, err := time.Parse(customDateFormat, dateStr)
	if err != nil {
		return fmt.Errorf("не удалось разобрать дату: %w", err)
	}
	cd.Time = parsedTime
	return nil
}

// MarshalJSON - кодирование CustomDate в JSON
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", cd.Time.Format(customDateFormat))), nil
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
	dateStr, ok := value.(string)
	if !ok {
		return fmt.Errorf("неподдерживаемый тип для CustomDate: %T", value)
	}
	parsedTime, err := time.Parse(customDateFormat, dateStr)
	if err != nil {
		return fmt.Errorf("не удалось разобрать дату: %w", err)
	}
	cd.Time = parsedTime
	return nil
}

type AgriWorkPlan struct {
	ID                     int       `gorm:"primaryKey;column:id" json:"id"`
	Status                 *string   `gorm:"column:status;size:255" json:"status"`
	WorkType               *string   `gorm:"column:work_type;size:255" json:"work_type"`
	WorkSubtype            *string   `gorm:"column:work_subtype;size:255" json:"work_subtype"`
	WorkTypeID             *int      `gorm:"column:work_type_id" json:"work_type_id"`
	Season                 *int      `gorm:"column:season" json:"season"`
	PlannedStartDate       CustomDate `gorm:"column:planned_start_date" json:"planned_start_date"`
	PlannedEndDate         CustomDate `gorm:"column:planned_end_date" json:"planned_end_date"`
	AdditionalInfo         *string   `gorm:"column:additional_info;size:255" json:"additional_info"`
	Description            *string   `gorm:"column:description;size:255" json:"description"`
	PlannedWaterRate       *float64  `gorm:"column:planned_water_rate" json:"planned_water_rate"`
	PlannedRowSpacing      *float64  `gorm:"column:planned_row_spacing" json:"planned_row_spacing"`
	PlannedPlantSpacing    *float64  `gorm:"column:planned_plant_spacing" json:"planned_plant_spacing"`
	PlannedDepth           *float64  `gorm:"column:planned_depth" json:"planned_depth"`
	PlannedSpeed           *float64  `gorm:"column:planned_speed" json:"planned_speed"`
	ResponsiblePersonID    *int      `gorm:"column:responsible_person_id" json:"responsible_person_id"`
	ExternalID             *string   `gorm:"column:external_id;size:255" json:"external_id"`
	GroupableID            *int      `gorm:"column:groupable_id" json:"groupable_id"`
	GroupableType          *string   `gorm:"column:groupable_type;size:255" json:"groupable_type"`
	AgroRecommendationID   *int      `gorm:"column:agro_recommendation_id" json:"agro_recommendation_id"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`
	NotifyResponsibleUsers *bool     `gorm:"column:notify_responsible_users;default:true" json:"notify_responsible_users"`
	CurrentSeasonID        *int      `gorm:"column:current_season_id" json:"current_season_id"`
}

func (AgriWorkPlan) TableName() string {
	return "agri_work_plans"
}
