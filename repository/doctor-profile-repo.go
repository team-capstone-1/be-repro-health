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
