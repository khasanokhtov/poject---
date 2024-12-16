package models

import "time"

// MaintenanceRecordRow - модель для таблицы maintenance_record_rows
type MaintenanceRecordRow struct {
	ID                  int       `json:"id" gorm:"primaryKey"`
	MaintenanceRecordID int       `json:"maintenance_record_id"`
	MaintenanceTypeID   int       `json:"maintenance_type_id"`
	Description         string    `json:"description"`    // Пустая строка, но не null
	RepairStage         string    `json:"repair_stage"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
