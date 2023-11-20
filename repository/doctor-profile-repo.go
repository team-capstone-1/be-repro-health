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
	var education []model.DoctorEducation

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&education)

	if tx.Error != nil {
		return education, tx.Error
	}

	return education, nil
}

func GetDoctorCertification(id uuid.UUID) ([]model.DoctorCertification, error) {
	var certification []model.DoctorCertification

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&certification)

	if tx.Error != nil {
		return certification, tx.Error
	}

	return certification, nil
}
