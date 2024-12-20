package services

import (
	"fmt"
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"integration-cropwise-v1/utils"
	"log"
	"strings"

	"gorm.io/gorm"
)

//Сохраняет данные компании в базу
func SaveCompany(email, token, name string) error {
	
	resetSearchPath := "SET search_path TO public"
    if err := database.DB.Exec(resetSearchPath).Error; err != nil {
        return fmt.Errorf("ошибка сброса search_path на public: %w", err)
    }


	// Приведение email к нормализованному виду
	email = strings.TrimSpace(strings.ToLower(email))

	// Проверка на существование компании
	var existingCompany models.Company
	err := database.DB.First(&existingCompany, "LOWER(TRIM(email)) = ?", email).Error
	if err == nil {
		return fmt.Errorf("компания с таким email уже существует")
	}

	//Получаем ID компании из Кропвайз
	companyID, err := FetchCompanyID(token)
	if err != nil {
		return fmt.Errorf("ошибка получения ID компании: %w", err)
	}

	 // Генерация имени схемы
	schemaName := strings.ToLower(strings.ReplaceAll(name, " ", "_"))

	// Сохраняем новую компанию
	company := models.Company{
		Email: email,
		Token: token,
		Name:  name,
		SchemaName: schemaName,
		CompanyID: companyID,
	}

	err = database.DB.Create(&company).Error
	if err != nil {
		return fmt.Errorf("ошибка сохранения компании: %w", err)
	}

	// Создаём схему
	err = CreateSchema(company.SchemaName)
	if err != nil {
		return fmt.Errorf("ошибка создания схемы для компании %s: %w", company.Name, err)
	}

	log.Printf("Компания %s успешно сохранена, и схема создана.", company.SchemaName)

	// Асинхронный запуск загрузки данных
	go func() {
		log.Printf("Запуск миграции и загрузки данных для компании %s...", company.Name)
		if err := utils.FetchDataForCompany(company); err != nil {
			log.Printf("Ошибка загрузки данных для компании %s: %v", company.Name, err)
		} else {
			log.Printf("Данные успешно загружены для компании %s.", company.Name)
		}
	}()

	return nil


}

func FindCompanyByEmail(email string) (*models.Company, error) {
	var company models.Company

	// Приведение email к нижнему регистру
	normalizedEmail := strings.TrimSpace(strings.ToLower(email))

	log.Printf("Запрос: SELECT * FROM companies WHERE LOWER(email) = '%s'", normalizedEmail)

	err := database.DB.First(&company, "LOWER(TRIM(email)) = ?", normalizedEmail).Error
	if err != nil {
		return nil, err
	}

	return &company, nil
}

// Проверка существования таблицы
func TableExists(db *gorm.DB, schemaName string, tableName string) (bool, error) {
    var exists bool
    query := fmt.Sprintf(
        "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s')",
        schemaName, tableName,
    )
    if err := db.Raw(query).Scan(&exists).Error; err != nil {
        log.Printf("Ошибка проверки таблицы %s.%s: %v", schemaName, tableName, err)
        return false, err
    }
    return exists, nil
}

func MigrateExistingSchemas() error {
    companies, err := database.GetAllCompanies()
    if err != nil {
        return fmt.Errorf("ошибка получения списка компаний: %w", err)
    }

    for _, company := range companies {
        schemaName := fmt.Sprintf(`"%s"`, company.SchemaName)

        // Устанавливаем search_path
        setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaName)
        if err := database.DB.Exec(setSearchPath).Error; err != nil {
            log.Printf("Ошибка переключения на схему %s: %v", schemaName, err)
            return err
        }

        // Логируем текущий search_path
        var currentSearchPath string
        if err := database.DB.Raw("SHOW search_path").Scan(&currentSearchPath).Error; err != nil {
            log.Printf("Ошибка получения search_path для схемы %s: %v", schemaName, err)
        } else {
            log.Printf("Текущий search_path для схемы %s: %s", schemaName, currentSearchPath)
        }

        log.Printf("Запускаем миграцию для схемы: %s", schemaName)

        // Выполняем миграцию
        if err := database.DB.AutoMigrate(
            &models.Machine{},
            &models.AdditionalObject{},
            &models.AgriWorkPlan{},
            &models.AgriWorkPlanApplicationMixItemModel{},
			&models.AgroOperation{},
			&models.ApplicationMixItem{},
			&models.Chemical{},
			&models.Crop{},
			&models.Fertilizer{},
			&models.FieldGroup{},
			&models.Field{},
			&models.FuelType{},
			&models.GroupFolder{},
			&models.HarvestWeighing{},
			&models.HistoryItem{},
			&models.Implement{},
			&models.MachineRegion{},
			&models.MachineRegionMappingItem{},
			&models.MachineGroup{},
			&models.MachineTaskAgriWorkPlanMappingItemModel{},
			&models.MachineTaskAgroOperationMappingItem{},
			&models.MachineTaskFieldMappingItem{},
			&models.MachineTask{},
			&models.User{},
			&models.WorkType{},
			&models.WorkTypeGroup{},
			); err != nil {
            log.Printf("Ошибка миграции для схемы %s: %v", schemaName, err)
            return err
        }

        log.Printf("Миграция завершена для схемы: %s", schemaName)
    }

    // Возвращаемся на public
    resetSearchPath := "SET search_path TO public"
    if err := database.DB.Exec(resetSearchPath).Error; err != nil {
        log.Printf("Ошибка сброса search_path на public: %v", err)
        return err
    }

    log.Println("Миграция для всех существующих схем завершена успешно.")
    return nil
}

// CreateSchema создаёт схему для компании и выполняет миграцию моделей
func CreateSchema(companyName string) error {
	schemaNameSQL := fmt.Sprintf(`"%s"`, companyName)

	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaNameSQL)
	if err := database.DB.Exec(query).Error; err != nil {
		log.Printf("Ошибка создания схемы %s: %v", schemaNameSQL, err)
		return err
	}

	setSearchPath := fmt.Sprintf("SET search_path TO %s", schemaNameSQL)
	if err := database.DB.Exec(setSearchPath).Error; err != nil {
		log.Printf("Ошибка переключения на схему %s: %v", schemaNameSQL, err)
		return err
	}

	// Логируем текущий search_path
	var currentSearchPath string
	if err := database.DB.Raw("SHOW search_path").Scan(&currentSearchPath).Error; err != nil {
		log.Printf("Ошибка получения search_path: %v", err)
	} else {
		log.Printf("Текущий search_path после переключения: %s", currentSearchPath)
	}

	tables := []struct {
		Model     interface{}
		TableName string
	}{
		{&models.Machine{}, "machine"},
		{&models.AdditionalObject{}, "additional_object"},
		{&models.AgriWorkPlan{}, "agri_work_plans"},
		{&models.AgriWorkPlanApplicationMixItemModel{}, "agri_work_plan_application_mix_item"},
		{&models.AgroOperation{}, "agro_operation"},
		{&models.ApplicationMixItem{}, "application_mix_item"},
		{&models.Chemical{}, "chemical"},
		{&models.Crop{}, "crop_model"},
		{&models.Fertilizer{}, "fertilizers"},
		{&models.FieldGroup{}, "field_groups"},
		{&models.Field{}, "fields"},
		{&models.FuelType{}, "fuel_type"},
		{&models.GroupFolder{}, "group_folders"},
		{&models.HarvestWeighing{}, "harvest_weighings"},
		{&models.HistoryItem{}, "history_items"},
		{&models.Implement{}, "implement"},
		{&models.MachineRegion{}, "machine_region"},
		{&models.MachineRegionMappingItem{}, "machine_region_mapping_items"},
		{&models.MachineGroup{}, "machine_groups"},
		{&models.MachineTaskAgriWorkPlanMappingItemModel{}, "machine_task_agri_work_plan_mapping_itens"},
		{&models.MachineTaskAgroOperationMappingItem{}, "machine_task_agro_operation_mapping_items"},
		{&models.MachineTaskFieldMappingItem{}, "machine_task_field_mapping_items"},
		{&models.MachineTask{}, "machine_task"},
		{&models.User{}, "users"},
		{&models.WorkType{}, "work_type"},
		{&models.WorkTypeGroup{}, "work_type_group"},
	}

	for _, table := range tables {
		exists, err := TableExists(database.DB, companyName, table.TableName)
		if err != nil {
			log.Printf("Ошибка проверки существования таблицы %s: %v", table.TableName, err)
			return err
		}
		if !exists {
			log.Printf("Таблица %s не существует. Выполняем миграцию.", table.TableName)
			if err := database.DB.AutoMigrate(table.Model); err != nil {
				log.Printf("Ошибка миграции таблицы %s: %v", table.TableName, err)
				return err
			}
		} else {
			log.Printf("Таблица %s уже существует. Миграция не требуется.", table.TableName)
		}
	}

	resetSearchPath := "SET search_path TO public"
	if err := database.DB.Exec(resetSearchPath).Error; err != nil {
		log.Printf("Ошибка сброса search_path на public: %v", err)
		return err
	}

	log.Printf("Схема %s успешно создана и миграция выполнена.", schemaNameSQL)
	return nil
}



