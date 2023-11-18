package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func GetAllClinics() ([]model.Clinic, error) {
	var dataclinics []model.Clinic

	tx := database.DB.Find(&dataclinics)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataclinics, nil
}