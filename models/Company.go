package models

type Company struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Email      string `gorm:"unique;not null"`
	Token      string `gorm:"not null"`
	Name       string `gorm:"not null"`
	SchemaName string `gorm:"unique;not null"`
	CompanyID  uint   `gorm:"unique;not null"`
}