package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func DoctorGetAllSchedules(doctorID uuid.UUID) ([]model.Consultation, error) {
	var consultation []model.Consultation

	tx := database.DB.Where("doctor_id = ?", doctorID).Preload("Patient")

	tx.Find(&consultation)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return consultation, nil
}
