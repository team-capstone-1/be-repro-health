package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func DoctorGetAppointment(name string, status string) ([]model.Consultation, error) {
	var dataConsultation []model.Consultation

	tx := database.DB.Model(&dataConsultation).Preload("Patient").Preload("Doctor").Preload("Clinic").Preload("Transaction").Preload("Transaction.Payment")

	if name != "" {
		tx = tx.Where("name LIKE ?", "%"+name+"%")
	}

	if status != "" {
		tx = tx.Where("status = ?", status)
	}

	if err := tx.Find(&dataConsultation).Error; err != nil {
		return nil, err
	}

	return dataConsultation, nil
}
