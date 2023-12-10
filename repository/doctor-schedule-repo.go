package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"fmt"

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

func GetPatientIDsByDateAndSession(doctorID uuid.UUID, session string) ([]uuid.UUID, error) {
	var consultations []model.Consultation

	// Implementasikan query untuk mendapatkan konsultasi yang sesuai
	err := database.DB.Where("doctor_id = ? AND session = ?", doctorID, session).
		Find(&consultations).Error

	if err != nil {
		return nil, err
	}

	// Ekstrak ID pasien dari hasil konsultasi
	var patientIDs []uuid.UUID
	for _, consultation := range consultations {
		patientIDs = append(patientIDs, consultation.PatientID)

	}
	fmt.Print("Ini adalah pasien id", patientIDs)

	return patientIDs, nil
}

func DoctorInactiveSchedule(doctorID uuid.UUID, date string, session string) (model.Consultation, error) {
	var doctorHoliday model.Consultation

	// Cari jadwal dokter pada tanggal dan sesi tertentu
	tx := database.DB.Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session)
	if tx.Error != nil {
		return doctorHoliday, tx.Error
	}

	// Ubah status doctor_available menjadi false
	tx = database.DB.Model(&doctorHoliday).Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).Update("doctor_available", false).Find(&doctorHoliday)
	if tx.Error != nil {
		return doctorHoliday, tx.Error
	}

	return doctorHoliday, nil
}

func UpdateTransactionStatusToWaiting(consultationID uuid.UUID) error {
	var consultations []model.Consultation

	// Find consultations based on the consultation ID
	tx := database.DB.Where("id = ?", consultationID).Find(&consultations)
	if tx.Error != nil {
		return tx.Error
	}

	// Update status to "waiting" for each associated transaction
	for _, consultation := range consultations {
		fmt.Printf("Processing Consultation ID: %s\n", consultation.ID)

		var transactions []model.Transaction
		tx := database.DB.Where("consultation_id = ?", consultation.ID).Find(&transactions)
		if tx.Error != nil {
			return tx.Error
		}

		for _, transaction := range transactions {
			fmt.Printf("Processing Transaction ID: %s\n", transaction.ID)

			tx := database.DB.Model(&model.Transaction{}).
				Where("id = ?", transaction.ID).
				Update("status", "waiting")

			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	return nil
}
