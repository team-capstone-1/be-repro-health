package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

func DoctorGetAllSchedules(doctorID uuid.UUID, session string, date time.Time) ([]model.Consultation, error) {
	var consultation []model.Consultation

	tx := database.DB.Model(&consultation).Where("doctor_id = ?", doctorID).Preload("Patient")

	if session != "" {
		tx = tx.Where("session = ?", session)
	}

	if !date.IsZero() {
		tx = tx.Where("date = ?", date)
	}

	tx.Find(&consultation)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return consultation, nil
}
