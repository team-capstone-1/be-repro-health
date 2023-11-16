package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func GetAllPatients() ([]model.Patient, error) {
	var datapatients []model.Patient

	tx := database.DB.Find(&datapatients)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datapatients, nil
}

func GetPatientByID(id int) (model.Patient, error) {
	var datapatient model.Patient

	tx := database.DB.First(&datapatient, id)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return datapatient, nil
}

func InsertPatient(data model.Patient) (model.Patient, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return data, nil
}

func DeletePatientByID(id int) error {
	tx := database.DB.Delete(&model.Patient{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdatePatientByID(id int, updateData model.Patient) (model.Patient, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return updateData, nil
}
