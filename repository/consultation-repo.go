package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func InsertConsultation(data model.Consultation) (model.Consultation, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}
	return data, nil
}