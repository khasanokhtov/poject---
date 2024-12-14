package database

import (
	"integration-cropwise-v1/models"
)


func GetAllCompanies()([]models.Company, error){
	var companies []models.Company
	if err := DB.Find(&companies).Error; err != nil{
		return nil, err
	}
	return companies, nil
}