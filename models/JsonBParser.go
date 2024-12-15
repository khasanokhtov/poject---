package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = JSONB{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		if err := json.Unmarshal(v, j); err != nil {
			return fmt.Errorf("ошибка парсинга JSONB: %w", err)
		}
	case string:
		if err := json.Unmarshal([]byte(v), j); err != nil {
			return fmt.Errorf("ошибка парсинга JSONB: %w", err)
		}
	default:
		return fmt.Errorf("неподдерживаемый тип: %T", value)
	}

	return nil
}
