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
            // &models.MachineModel{},
            // &models.AdditionalObjectModel{},
            // &models.AgriWorkPlanModel{},
            // &models.AgriWorkPlanApplicationMixItemModel{},
			// &models.AgronomistAssignmentModel{},
			// &models.AgronomistFieldModel{},
			// &models.AgroOperationModel{},
			// &models.ApplicationMixItemModel{},
			// &models.AllowedToCropModel{},
			// &models.AvatarModel{},
			// &models.ChemicalModel{},
			// &models.CounterpartyModel{},
			// &models.CropModel{},
			// &models.DataSourceGPSLoggerModel{},
			// &models.DataSourceParameterModel{},
			// &models.EquipmentAssignmentModel{},
			// &models.FertilizerModel{},
			// &models.FieldGroupModel{},
			// &models.FieldScoutReportModel{},
			// &models.FuelHourlyDataItemModel{},
			// &models.FieldShapeModel{},
			// &models.FieldShapeLandParcelMappingItemModel{},
			// &models.FieldModel{},
			// &models.FuelTypeModel{},
			// &models.GPSLoggerModel{},
			// &models.GPSLoggerMappingItemModel{},
			// &models.GroupFolderModel{},
			// &models.HarvestIndicatorModel{},
			// &models.HarvestWeighingModel{},
			// &models.HistoryItemModel{},
			// &models.ImplementModel{},
			// &models.ImplementRegionMappingItemModel{},
			// &models.ImplementWorkTypeMappingItemModel{},
			// &models.InventoryHistoryItemModel{},
			// &models.MachineDowntimeModel{},
			// &models.MachineDowntimeTypeModel{},
			// &models.MachineDowntimeTypeGroupModel{},
			// &models.MachineRegionModel{},
			// &models.MachineRegionMappingItemModel{},
			// &models.MachineGroupModel{},
			// &models.MachineObjectIntersectionModel{},
			// &models.MachineTaskAgriWorkPlanMappingItemModel{},
			// &models.MachineTaskAgroOperationMappingItemModel{},
			// &models.MachineTaskFieldMappingItemModel{},
			// &models.MachineTaskModel{},
			// &models.MachineTaskGroupFolderMappingItemModel{},
			// &models.MachineWorkPlanModel{},
			// &models.MachineWorkPlanRowModel{},
			// &models.MachineWorkPlanItemModel{},
			// &models.MachineryManufacturerModel{},
			// &models.MachineryModel{},
			// &models.MaintenanceType{},
			// &models.MaintenanceTypeGroup{},
			// &models.MaintenanceRecord{},
			// &models.MaintenanceRecordRow{},
			// &models.OdometerState{},
			// &models.PersonalIdentifier{},
			// &models.ProductionCycle{},
			// &models.ProductivityEstimate{},
			// &models.ProductivityEstimateHistory{},
			// &models.Users{},
			//&models.WorkType{},
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
		// {&models.MachineModel{}, "machine"},
		// {&models.AdditionalObjectModel{}, "additional_object"},
		// {&models.AgriWorkPlanModel{}, "agri_work_plans"},
		// {&models.AgriWorkPlanApplicationMixItemModel{}, "agri_work_plan_application_mix_item"},
		// {&models.AgronomistAssignmentModel{}, "agronomist"},
		// {&models.AgronomistFieldModel{}, "agronomist_field"},
		// {&models.AgroOperationModel{}, "agro_operation"},
		// {&models.ApplicationMixItemModel{}, "application_mix_item"},
		// {&models.AllowedToCropModel{}, "allowed_to_crop"},
		// {&models.AvatarModel{}, "avatar"},
		// {&models.ChemicalModel{}, "chemical"},
		// {&models.CounterpartyModel{}, "counterparty_models"},
		// {&models.CropModel{}, "crop_model"},
		// {&models.DataSourceGPSLoggerModel{}, "data_source_gps_logger"},
		// {&models.DataSourceParameterModel{}, "data_source_parametr"},
		// {&models.EquipmentAssignmentModel{}, "equipment_assignment"},
		// {&models.FertilizerModel{}, "fertilizers"},
		// {&models.FieldGroupModel{}, "field_groups"},
		// {&models.FieldScoutReportModel{}, "field_scout_reports"},
		// {&models.FuelHourlyDataItemModel{}, "fuel_hourly_data_item"},
		// {&models.FieldShapeModel{}, "field_shapes"},
		// {&models.FieldShapeLandParcelMappingItemModel{}, "field_shapes_land_parcels"},
		// {&models.FieldModel{}, "fields"},
		// {&models.FuelTypeModel{}, "fuel_type"},
		// {&models.GPSLoggerModel{}, "gps_loggers"},
		// {&models.GPSLoggerMappingItemModel{}, "gps_loggers_mapping_item"},
		// {&models.GroupFolderModel{}, "group_folders"},
		// {&models.HarvestIndicatorModel{}, "harvest_indicators"},
		// {&models.HarvestWeighingModel{}, "harvest_weighings"},
		// {&models.HistoryItemModel{}, "history_items"},
		// {&models.ImplementModel{}, "implement"},
		// {&models.ImplementRegionMappingItemModel{}, "implement_region_mapping_items"},
		// {&models.ImplementWorkTypeMappingItemModel{}, "implement_work_type_mapping_items"},
		// {&models.InventoryHistoryItemModel{}, "inventory_history_items"},
		// {&models.MachineDowntimeModel{}, "machine_downtimes"},
		// {&models.MachineDowntimeTypeModel{}, "machine_downtime_type"},
		// {&models.MachineDowntimeTypeGroupModel{}, "machine_downtime_type_group"},
		// {&models.MachineRegionModel{}, "machine_region"},
		// {&models.MachineRegionMappingItemModel{}, "machine_region_mapping_items"},
		// {&models.MachineGroupModel{}, "machine_groups"},
		// {&models.MachineObjectIntersectionModel{}, "machine_object_intersections"},
		// {&models.MachineTaskAgriWorkPlanMappingItemModel{}, "machine_task_agri_work_plan_mapping_itens"},
		// {&models.MachineTaskAgroOperationMappingItemModel{}, "machine_task_agro_operation_mapping_items"},
		// {&models.MachineTaskFieldMappingItemModel{}, "machine_task_field_mapping_items"},
		// {&models.MachineTaskModel{}, "machine_task"},
		// {&models.MachineTaskGroupFolderMappingItemModel{}, "machine_task_group_folder_mapping_items"},
		// {&models.MachineWorkPlanModel{}, "machine_work_plans"},
		// {&models.MachineWorkPlanRowModel{}, "machine_work_plans_row"},
		// {&models.MachineWorkPlanItemModel{}, "machine_work_plan_items"},
		// {&models.MachineryManufacturerModel{}, "machine_manufacturer"},
		// {&models.MachineryModel{}, "machinery_model"},
		// {&models.MaintenanceType{}, "maintence_type"},
		// {&models.MaintenanceTypeGroup{}, "maintence_type_group"},
		// {&models.MaintenanceRecord{}, "maintence_record"},
		// {&models.MaintenanceRecordRow{}, "maintence_recor_row"},
		// {&models.OdometerState{}, "odomentr_state"},
		// {&models.PersonalIdentifier{}, "personal_identifiers"},
		// {&models.ProductionCycle{}, "production_cycle"},
		// {&models.ProductivityEstimate{}, "productivity_estimate"},
		// {&models.ProductivityEstimateHistory{}, "estimate_history"},
		// {&models.Users{}, "users"},
		// {&models.WorkType{}, "work_type"},
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



