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
	tx := database.DB.Model(&data).Where("consultation_id = ?", consultationID).Update("status", "processed")
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DoctorFinishConsultation(consultationID uuid.UUID, data model.Transaction) error {
	tx := database.DB.Model(&data).Where("consultation_id = ?", consultationID).Update("status", "done")
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

func DoctorGetDetailsConsultation(consultationID uuid.UUID) (model.Consultation, error) {
	var consultation model.Consultation

	tx := database.DB.
		Where("id = ?", consultationID).
		Preload("Patient").
		Preload("Patient.User").
		Preload("Clinic").
		First(&consultation)

	if tx.Error != nil {
		return consultation, tx.Error
	}

	return consultation, nil
}

func DoctorGetTransactionsForConsultation(consultationID uuid.UUID) ([]model.Transaction, error) {
	var transactions []model.Transaction

	tx := database.DB.
		Where("consultation_id = ?", consultationID).
		Preload("Payment").
		Find(&transactions)

	if tx.Error != nil {
		return transactions, tx.Error
	}

	return transactions, nil
}

func DoctorGetAllConsultations(doctorID uuid.UUID, name string, status string) ([]model.Consultation, error) {
	var consultation []model.Consultation

	tx := database.DB.
		Joins("JOIN patients ON patients.id = consultations.patient_id").
		Joins("JOIN transactions ON transactions.consultation_id = consultations.id").
		Preload("Patient").
		Preload("Transaction.Payment").
		Where("consultations.doctor_id = ?", doctorID)

	if name != "" {
		tx = tx.Where("patients.name LIKE ?", "%"+name+"%")
	}

	if status != "" {
		tx = tx.Where("transactions.status = ?", status)
	}

	tx = tx.Find(&consultation)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return consultation, nil
}
