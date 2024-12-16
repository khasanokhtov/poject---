package models

import "time"

// MachineTaskGroupFolderMappingItemModel - модель для таблицы machine_task_group_folder_mapping_items
type MachineTaskGroupFolderMappingItemModel struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	GroupFolderID int       `json:"group_folder_id"`
	MachineTaskID int       `json:"machine_task_id"`
	ExternalID   *string    `json:"external_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
