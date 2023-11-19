package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetDoctorProfile(id uuid.UUID) (model.Doctor, error) {
	var profileDoctor model.Doctor

	tx := database.DB.Preload("Specialist").Preload("Clinic").Where("id = ?", id).First(&profileDoctor)

	if tx.Error != nil {
		return profileDoctor, tx.Error
	}

	return profileDoctor, nil
}

func GetDoctorWorkHistory(id uuid.UUID) ([]model.DoctorWorkHistory, error) {
	var workHistory []model.DoctorWorkHistory

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&workHistory)

	if tx.Error != nil {
		return workHistory, tx.Error
	}

	return workHistory, nil
}

func GetDoctorEducation(id uuid.UUID) ([]model.DoctorEducation, error) {
	var doctorEducation []model.DoctorEducation

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&doctorEducation)

	if tx.Error != nil {
		return doctorEducation, tx.Error
	}

	return doctorEducation, nil
}
