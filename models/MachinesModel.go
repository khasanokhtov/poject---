package models

type MachineResponse struct {
	Data []MachineModel `json:"data"`
	Meta struct {
		Request struct {
			FromID     int    `json:"from_id"`
			Limit      *int   `json:"limit"`
			ServerTime string `json:"server_time"`
		} `json:"request"`
		Response struct {
			Limit           int `json:"limit"`
			ObtainedRecords int `json:"obtained_records"`
			FirstRecordID   int `json:"first_record_id"`
			LastRecordID    int `json:"last_record_id"`
		} `json:"response"`
	} `json:"meta"`
}

type MachineModel struct {
	ID                        int      `json:"id"`
	Name                      string   `json:"name"`
	Model                     string   `json:"model"`
	Manufacturer              string   `json:"manufacturer"`
	Year                      *int     `json:"year"` // Может быть null
	RegistrationNumber        string   `json:"registration_number"`
	InventoryNumber           string   `json:"inventory_number"`
	MachineGroupID            int      `json:"machine_group_id"`
	MachineType               string   `json:"machine_type"`
	MachineSubtype            string   `json:"machine_subtype"`
	AvatarID                  int      `json:"avatar_id"`
	ChassisSerialNumber       string   `json:"chassis_serial_number"`
	EngineSerialNumber        string   `json:"engine_serial_number"`
	EnginePower               *float64 `json:"engine_power"` // Может быть null
	FuelType                  string   `json:"fuel_type"`
	FuelTankSize              *float64 `json:"fuel_tank_size"`        // Может быть null
	FuelConsumptionNorm       *float64 `json:"fuel_consumption_norm"` // Может быть null
	LegalCompany              string   `json:"legal_company"`
	Description               string   `json:"description"`
	DefaultImplementID        *int     `json:"default_implement_id"` // Может быть null
	DefaultDriverID           int      `json:"default_driver_id"`
	Additional1               *string  `json:"additional_1"` // Может быть null
	Additional2               *string  `json:"additional_2"` // Может быть null
	AdditionalInfo            string   `json:"additional_info"`
	PhoneNumber               string   `json:"phone_number"`
	CreatedAt                 string   `json:"created_at"`
	UpdatedAt                 string   `json:"updated_at"`
	ExternalID                *string  `json:"external_id"` // Может быть null
	MachineryModelID          int      `json:"machinery_model_id"`
	FuelTypeID                int      `json:"fuel_type_id"`
	RefuelSource              string   `json:"refuel_source"`
	MachineryManufacturerID   int      `json:"machinery_manufacturer_id"`
	EngineCapacity            *float64 `json:"engine_capacity"` // Может быть null
	Weight                    *float64 `json:"weight"`          // Может быть null
	Height                    *float64 `json:"height"`          // Может быть null
	Width                     *float64 `json:"width"`           // Может быть null
	Length                    *float64 `json:"length"`          // Может быть null
	UnchangedDefaultImplement bool     `json:"unchanged_default_implement"`
	MinDowntimeInSeconds      int      `json:"min_downtime_in_seconds"`
	CalculateDowntimes        string   `json:"calculate_downtimes"`
}
