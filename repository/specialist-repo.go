package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func GetAllSpecialists() ([]model.Specialist, error) {
	var dataspecialists []model.Specialist

	tx := database.DB.Find(&dataspecialists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataspecialists, nil
}