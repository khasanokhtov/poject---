package main

import (
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/routes"
	"integration-cropwise-v1/services"
	"integration-cropwise-v1/utils"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)



func main() {
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла для логов: %v", err)
	}
	defer logFile.Close()

	// Установка вывода логов в файл
	log.SetOutput(logFile)
	log.Println("Логирование начато")
	
	//DB Connect
	database.ConnectDB()
	app := fiber.New()
	routes.SetupAuthRoutes(app)
	routes.SetupCompanyRoutes(app)
	log.Fatal(app.Listen(":3000"))

	 // Миграция существующих схем
	if err := services.MigrateExistingSchemas(); err != nil {
        log.Fatalf("Ошибка миграции существующих схем: %v", err)
    }

	
	
	//Запуск worker-пула
	go func() {
		for {
			log.Println("Начало ежедневного обновления данных.")
			utils.FetchingData()
			log.Println("Ежедневное обновление данных завершено.")
			time.Sleep(24 * time.Hour) // Задержка на 24 часа
		}
	}()
	log.Println("Все компании успешно обработаны.")
	
}