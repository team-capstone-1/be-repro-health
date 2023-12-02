package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
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

func DoctorConfirmConsultation(consultationID uuid.UUID, data model.Transaction) error {
	tx := database.DB.Where("id = ?", consultationID).Update("status", "processed")
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DoctorFinishConsultation(consultationID uuid.UUID, data model.Transaction) error {
	tx := database.DB.Where("id = ?", consultationID).Update("status", "done")
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DoctorGetDetailsTransaction(transactionID uuid.UUID) (model.Transaction, error) {
	var transaction model.Transaction
	tx := database.DB.Where("id = ?", transactionID).Preload("Payment").First(&transaction)

	if tx.Error != nil {
		return transaction, tx.Error
	}

	return transaction, nil
}
