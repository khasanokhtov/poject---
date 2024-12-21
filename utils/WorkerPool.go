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
		if err := tasks.FetchAndSaveMachines(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки машин для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAdditionalObjects(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки дополнительных объектов для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveChemicals(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки дополнительных объектов для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgriWorkPlans(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки агроопераций для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgriWorkPlanApplicationMixItems(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки agriworkplan_application_mix_items для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgroOperations(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии agro_operations компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveApplicationMixItems(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки application_mix_items для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveCrops(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки crops для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveFertilizers(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки fertilizers для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveFieldGroups(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки field_groups для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveFields(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки fields для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveFuelTypes(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки fuel_type для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveGroupFolders(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки group_folders для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveHarvestWeighings(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки harvest_weighings для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveHistoryItems(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки history_items для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveImplements(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки implemets для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineRegions(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_regions для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineRegionMappingItems(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_region_mapping_item для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineGroups(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_groups для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineTaskAgriWorkPlanMappingItems(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_task_agri_work_plan_mapping_itens для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineTaskAgroOperationMappingItems(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_task_agro_operation_mapping_items для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineTaskFieldMappingItems(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_task_agro_operation_mapping_items для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveMachineTasks(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки machine_task для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveUsers(database.DB,company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки users для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveWorkTypes(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки work_type для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveWorkTypeGroups(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки work_type_groups для компании %s: %v", company.Name, err)
			continue
		}
	}
	return nil
}

