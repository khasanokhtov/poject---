package utils

import (
	"integration-cropwise-v1/database"
	"integration-cropwise-v1/models"
	"integration-cropwise-v1/tasks"

	"log"
)

func FetchingData() {
	companies, err := database.GetAllCompanies()
	if err != nil {
		log.Fatalf("Ошибка получения списка компаний: %v", err)
	}

	for _, company := range companies {
		if err := FetchDataForCompany(company); err != nil {
			log.Printf("Ошибка обработки компании %s: %v", company.Name, err)
		}
	}
}


func FetchDataForCompany(company models.Company) error{

	companies, err := database.GetAllCompanies()
	if err != nil {
		log.Fatalf("Ошибка получения списка компаний: %v", err)
	}

	for _, company := range companies {
		log.Printf("Обрабатываем компанию: %s", company.Name)
		// if err := tasks.FetchAndSaveMachines(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки машин для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAdditionalObjects(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки дополнительных объектов для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAgriWorkPlans(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки агроопераций для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAgriWorkPlanApplicationMixItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки agriworkplan_application_mix_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAgronomistAssignments(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки agronomist_assignments для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAgronomistFields(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки agronomist_fields для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAgroOperations(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки химии agro_operations компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveApplicationMixItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки application_mix_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAllowedToCrops(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки allowed_to_crops для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveAvatars(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки avatars для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveChemicals(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveCounterparties(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки counterpairties для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveCrops(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки crops для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveDataSourceGPSLoggers(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки GPSloggers для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveDataSourceParameters(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки data_source_parametr для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveEquipmentAssignments(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки equipments_assignment для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFertilizers(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки fertilizers для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFieldGroups(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки field_groups для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFieldScoutReports(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки field_scout_reports для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFuelHourlyDataItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки fuel_hourly_data_item для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFieldShapes(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки field_shapes для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFieldShapeLandParcelMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки field_shapes_land_parcels для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFields(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки fields для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveFuelTypes(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки fuel_type для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveGPSLoggers(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки gps_loggers для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveGPSLoggerMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки gps_loggers_mapping_item для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveGroupFolders(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки group_folders для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveHarvestIndicators(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки harvest_indicators для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveHarvestWeighings(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки harvest_weighings для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveHistoryItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки history_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveImplements(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки implemets для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveImplementRegionMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки implement_region_mapping_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveImplementWorkTypeMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки implement_work_type_mapping_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveInventoryHistoryItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки inventory_history_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineDowntimes(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_downtimes для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineDowntimeTypes(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_downtime_type для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineDowntimeTypeGroups(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_downtime_type_group для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineRegions(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_regions для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineRegionMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_region_mapping_item для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineGroups(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_groups для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineObjectIntersections(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_object_intersections для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineTaskAgriWorkPlanMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_task_agri_work_plan_mapping_itens для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineTaskAgroOperationMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_task_agro_operation_mapping_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineTaskFieldMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_task_agro_operation_mapping_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineTasks(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_task для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineTaskGroupFolderMappingItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_task_group_folder_mapping_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineWorkPlans(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_work_plans для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineWorkPlanRows(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_work_plans_row для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineWorkPlanItems(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_work_plan_items для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineryManufacturers(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_manufacturers для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMachineryModels(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки machine_models для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMaintenanceTypes(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки maintence_type для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMaintenanceTypeGroups(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки maintence_type_groups для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMaintenanceRecords(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки maintence_record для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveMaintenanceRecordRows(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки maintence_record_rows для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveOdometerStates(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки odometr_state для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSavePersonalIdentifiers(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки personal_identifiers для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveProductionCycles(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки production_cycle для компании %s: %v", company.Name, err)
		// 	continue
		// }
		// if err := tasks.FetchAndSaveProductivityEstimates(company.Token, company.SchemaName); err != nil {
		// 	log.Printf("Ошибка загрузки productivity_estimate для компании %s: %v", company.Name, err)
		// 	continue
		// }
		if err := tasks.FetchAndSaveProductivityEstimateHistories(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки productivity_estimate_history для компании %s: %v", company.Name, err)
			continue
		}
	}
	return nil
}

