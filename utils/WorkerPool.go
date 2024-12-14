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
		if err := tasks.FetchAndSaveMachines(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки машин для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAdditionalObjects(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки дополнительных объектов для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgriWorkPlans(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки агроопераций для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgriWorkPlanApplicationMixItems(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgronomistAssignments(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgronomistFields(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAgroOperations(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveApplicationMixItems(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAllowedToCrops(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveAvatars(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveChemicals(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
		if err := tasks.FetchAndSaveCounterparties(company.Token, company.SchemaName); err != nil {
			log.Printf("Ошибка загрузки химии для компании %s: %v", company.Name, err)
			continue
		}
	}
	return nil
}