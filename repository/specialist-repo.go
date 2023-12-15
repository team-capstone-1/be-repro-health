package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetSpecialists() ([]model.Specialist, error) {
	var specialists []model.Specialist

	tx := database.DB.Find(&specialists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return specialists, nil
}

func GetSpecialistByID(id uuid.UUID) (model.Specialist, error) {
	var specialist model.Specialist

	tx := database.DB.First(&specialist, id)
	if tx.Error != nil {
		return specialist, tx.Error
	}

	return specialist, nil
}

func GetSpecialistsByClinic(id uuid.UUID) ([]model.Specialist, error) {
	var clinic model.Clinic

	tx := database.DB.Preload("Doctors.Specialist").Find(&clinic, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Use a map to store unique specialists
	uniqueSpecialists := make(map[uuid.UUID]model.Specialist)
	for _, doctor := range clinic.Doctors {
		uniqueSpecialists[doctor.Specialist.ID] = doctor.Specialist
	}

	// Convert the map values to a slice
	var specialists []model.Specialist
	for _, specialist := range uniqueSpecialists {
		specialists = append(specialists, specialist)
	}

	return specialists, nil
}


func InsertSpecialist(specialist model.Specialist) (model.Specialist, error) {
tx := database.DB.Create(&specialist)
	if tx.Error != nil {
		return model.Specialist{}, tx.Error
	}
	return specialist, nil
}

func DeleteSpecialistByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Specialist{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateSpecialistDoctorByID(id uuid.UUID, updateData model.Specialist) (model.Specialist, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(&updateData)
	if tx.Error != nil {
		return model.Specialist{}, tx.Error
	}
	return updateData, nil
}
