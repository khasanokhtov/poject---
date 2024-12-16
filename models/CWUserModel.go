package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// AvatarForUsers - структура для вложенного объекта avatar
type AvatarForUsers struct {
	URL          string  `json:"url"`
	AvatarLarge  *Image  `json:"avatar_large"` // Может быть null
	AvatarMid    *Image  `json:"avatar_mid"`   // Может быть null
	AvatarSmall  *Image  `json:"avatar_small"` // Может быть null
	AvatarTiny   *Image  `json:"avatar_tiny"`  // Может быть null
}

// Image - структура для вложенного объекта изображений
type Image struct {
	URL string `json:"url"`
}

// Реализуем интерфейсы driver.Valuer и sql.Scanner для AvatarForUsers
func (a AvatarForUsers) Value() (driver.Value, error) {
	if (a == AvatarForUsers{}) {
		return nil, nil
	}
	return json.Marshal(a)
}

func (a *AvatarForUsers) Scan(value interface{}) error {
	if value == nil {
		*a = AvatarForUsers{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("не удалось преобразовать %T в AvatarForUsers", value)
	}
	return json.Unmarshal(bytes, a)
}

// UnitsTable - структура для вложенного объекта units_table
type UnitsTable struct {
	Area                   string `json:"area"`
	Depth                  string `json:"depth"`
	Speed                  string `json:"speed"`
	Length                 string `json:"length"`
	Volume                 string `json:"volume"`
	Weight                 string `json:"weight"`
	Pressure               string `json:"pressure"`
	WaterRate              string `json:"water_rate"`
	WindSpeed              string `json:"wind_speed"`
	YieldMass              string `json:"yield_mass"`
	RowSpacing             string `json:"row_spacing"`
	TankVolume             string `json:"tank_volume"`
	Temperature            string `json:"temperature"`
	AppliedMass            string `json:"applied_mass"`
	Productivity           string `json:"productivity"`
	ShortLength            string `json:"short_length"`
	PlantSpacing           string `json:"plant_spacing"`
	YieldDensity           string `json:"yield_density"`
	AppliedVolume          string `json:"applied_volume"`
	EngineCapacity         string `json:"engine_capacity"`
	FuelConsumption        string `json:"fuel_consumption"`
	MachineryWeight        string `json:"machinery_weight"`
	YieldAreaDensity       string `json:"yield_area_density"`
	AppliedVolumeRate      string `json:"applied_volume_rate"`
	PrecipitationLevel     string `json:"precipitation_level"`
	AppliedAreaDensity     string `json:"applied_area_density"`
	AppliedCntPerArea      string `json:"applied_cnt_per_area"`
	MachineryDimensions    string `json:"machinery_dimensions"`
	FuelConsumptionPerArea string `json:"fuel_consumption_per_area"`
	FuelConsumptionPerTime string `json:"fuel_consumption_per_time"`
}

// Users - модель для таблицы users
type Users struct {
	ID                           int             `json:"id" gorm:"primaryKey"`
	Username                     string          `json:"username"`
	Email                        string          `json:"email"`
	NotificationEmail            *string         `json:"notification_email"`       // Может быть null
	MobilePhone                  string          `json:"mobile_phone"`
	Avatar                       *AvatarForUsers `json:"avatar" gorm:"type:jsonb"` // Сохраняем как JSON
	Position                     *string         `json:"position"`                 // Может быть null
	Language                     string          `json:"language"`
	TimeZone                     string          `json:"time_zone"`
	YieldUnits                   string          `json:"yield_units"`
	Status                       string          `json:"status"`
	Dispatcher                   bool            `json:"dispatcher"`
	Driver                       bool            `json:"driver"`
	AdditionalInfo               *string         `json:"additional_info"`          // Может быть null
	Description                  *string         `json:"description"`              // Может быть null
	UnitsTable                   UnitsTable      `json:"units_table" gorm:"embedded;embeddedPrefix:units_table_"`
	RFID                         string          `json:"rfid"`
	LastSignInAt                 *time.Time      `json:"last_sign_in_at"`          // Может быть null
	CurrentSignInAt              *time.Time      `json:"current_sign_in_at"`       // Может быть null
	ConsultingCompanyID          *int            `json:"consulting_company_id"`    // Может быть null
	CreatedAt                    time.Time       `json:"created_at"`
	UpdatedAt                    time.Time       `json:"updated_at"`
	ExternalID                   *string         `json:"external_id"`              // Может быть null
	AuthMethod                   string          `json:"auth_method"`
	CreateScoutReportOutsideOfField bool         `json:"create_scout_report_outside_of_field"`
	Supervisor                   bool            `json:"supervisor"`
	Worker                       bool            `json:"worker"`
}
