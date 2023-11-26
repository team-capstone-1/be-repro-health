package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetAllPatients(user uuid.UUID) ([]model.Patient, error) {
	var dataPatients []model.Patient

	tx := database.DB.Where("user_id = ?", user).Find(&dataPatients)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return dataPatients, nil
}

func GetAllPatientsDashboard(doctorID uuid.UUID) ([]model.Patient, error) {
	var dataPatients []model.Patient

	tx := database.DB.Where("user_id = ?", doctorID).Find(&dataPatients)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return dataPatients, nil
}

func GetPatientByID(id uuid.UUID) (model.Patient, error) {
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

func DeletePatientByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Patient{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdatePatientByID(id uuid.UUID, updateData model.Patient) (model.Patient, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return updateData, nil
}
