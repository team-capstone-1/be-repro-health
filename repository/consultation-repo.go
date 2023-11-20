package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func GetAllConsultation() ([]model.Consultation, error) {
	var dataconsultations []model.Consultation

	tx := database.DB.Find(&dataconsultations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataconsultations, nil
}

func InsertConsultation(data model.Consultation) (model.Consultation, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}
	return data, nil
}