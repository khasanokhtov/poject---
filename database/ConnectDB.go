package database

import (
	"integration-cropwise-v1/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123654 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	DB = DB.Debug()

	log.Println("Подключение к базе данных установлено!")


	err = DB.AutoMigrate(
		&models.Company{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}
}