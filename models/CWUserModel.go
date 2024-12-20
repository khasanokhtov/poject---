package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Avatar struct {
	URL          string  `json:"url"`
	AvatarLarge  *Image  `json:"avatar_large"`
	AvatarMid    *Image  `json:"avatar_mid"`
	AvatarSmall  *Image  `json:"avatar_small"`
	AvatarTiny   *Image  `json:"avatar_tiny"`
}

type Image struct {
	URL string `json:"url"`
}

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

type User struct {
	ID                           int             `gorm:"primaryKey;column:id"`
	Username                     string          `gorm:"column:username;size:255"`
	Email                        string          `gorm:"column:email;size:255"`
	NotificationEmail            *string         `gorm:"column:notification_email;size:255"`
	MobilePhone                  string          `gorm:"column:mobile_phone;size:255"`
	Avatar                       *Avatar         `gorm:"column:avatar;type:jsonb"`
	Position                     *string         `gorm:"column:position;size:255"`
	Language                     string          `gorm:"column:language;size:255"`
	TimeZone                     string          `gorm:"column:time_zone;size:255"`
	YieldUnits                   string          `gorm:"column:yield_units;size:255"`
	Status                       string          `gorm:"column:status;size:255"`
	Dispatcher                   *bool           `gorm:"column:dispatcher"`
	Driver                       *bool           `gorm:"column:driver"`
	AdditionalInfo               *string         `gorm:"column:additional_info;type:text"`
	Description                  *string         `gorm:"column:description;type:text"`
	RFID                         string          `gorm:"column:rfid;size:255"`
	LastSignInAt                 *time.Time      `gorm:"column:last_sign_in_at"`
	CurrentSignInAt              *time.Time      `gorm:"column:current_sign_in_at"`
	ConsultingCompanyID          *int            `gorm:"column:consulting_company_id"`
	CreatedAt                    time.Time       `gorm:"column:created_at"`
	UpdatedAt                    time.Time       `gorm:"column:updated_at"`
	ExternalID                   *string         `gorm:"column:external_id;size:255"`
	AuthMethod                   string          `gorm:"column:auth_method;size:255"`
	CreateScoutReportOutsideOfField *bool        `gorm:"column:create_scout_report_outside_of_field"`
}

func (User) TableName() string {
	return "users"
}

func (a Avatar) Value() (driver.Value, error) {
	if (a == Avatar{}) {
		return nil, nil
	}
	return json.Marshal(a)
}

func (a *Avatar) Scan(value interface{}) error {
	if value == nil {
		*a = Avatar{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("не удалось преобразовать %T в Avatar", value)
	}
	return json.Unmarshal(bytes, a)
}