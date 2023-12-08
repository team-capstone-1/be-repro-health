package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func DoctorGetAllSchedules(doctorID uuid.UUID, session string, date string) ([]model.Consultation, error) {
	var consultation []model.Consultation

	tx := database.DB.Model(&consultation).Where("doctor_id = ?", doctorID).Preload("Patient")

	if session != "" {
		tx = tx.Where("session = ?", session)
	}

	if date != "" {
		tx = tx.Where("date = ?", date)
	}

	if err := tx.Find(&consultation).Error; err != nil {
		return nil, err
	}
	return consultation, nil
}
