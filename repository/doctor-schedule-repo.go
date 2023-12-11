package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func GetPatientIDsByDateAndSession(doctorID uuid.UUID, date string, session string) ([]uuid.UUID, error) {
	var consultations []model.Consultation

	// Implementasikan query untuk mendapatkan konsultasi yang sesuai
	err := database.DB.Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).
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

	// Cek apakah sudah ada record di DoctorHoliday
	exists, err := IsDoctorHolidayExists(doctorID, date, session)
	if err != nil {
		return doctorHoliday, err
	}

	if !exists {
		// Jika tidak ada, ubah status doctor_available menjadi false
		tx := database.DB.Model(&model.Consultation{}).
			Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).
			Update("doctor_available", false).
			Find(&doctorHoliday)

		if tx.Error != nil {
			return doctorHoliday, tx.Error
		}

		// Buat rekaman baru di DoctorHoliday
		err := CreateDoctorHoliday(doctorID, date, session)
		if err != nil {
			return doctorHoliday, err
		}
	}

	return doctorHoliday, nil
}

func IsDoctorHolidayExists(doctorID uuid.UUID, date, session string) (bool, error) {
	var count int64
	err := database.DB.Model(&model.DoctorHoliday{}).
		Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func UpdateTransactionStatusToWaiting(dateString, session string) error {

	// Find consultations based on the date and session
	var consultations []model.Consultation
	tx := database.DB.
		Where("date = ? AND session = ?", dateString, session).
		Find(&consultations)

	if tx.Error != nil {
		return tx.Error
	}

	// Iterate over the consultations and update associated transactions
	for _, consultation := range consultations {
		fmt.Printf("Processing Consultation ID: %s\n", consultation.ID)

		// Find transactions associated with the consultation
		var transactions []model.Transaction
		tx := database.DB.
			Where("consultation_id = ?", consultation.ID).
			Find(&transactions)

		if tx.Error != nil {
			return tx.Error
		}

		// Update status to "waiting" for each associated transaction
		for _, transaction := range transactions {
			fmt.Printf("Processing Transaction ID: %s\n", transaction.ID)

			// Update the status to "waiting"
			tx := database.DB.Model(&transaction).
				Update("status", "waiting")

			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	return nil
}

func CreateDoctorHoliday(doctorID uuid.UUID, date string, session string) error {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	doctorHoliday := model.DoctorHoliday{
		ID:       uuid.New(),
		DoctorID: doctorID,
		Date:     parsedDate,
		Session:  session,
	}

	if err := database.DB.Create(&doctorHoliday).Error; err != nil {
		return err
	}

	return nil
}

func IsDoctorOnHoliday(doctorID uuid.UUID, date string, session string) (bool, error) {
	var doctorHoliday model.DoctorHoliday
	err := database.DB.
		Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).
		First(&doctorHoliday).Error

	if err == gorm.ErrRecordNotFound {
		// Record tidak ditemukan, artinya dokter tidak sedang libur
		fmt.Println("Doctor is not on holiday.")
		return false, nil
	} else if err != nil {
		// Terjadi kesalahan lain
		fmt.Println("Error checking doctor's holiday:", err)
		return false, err
	}

	// Dokter sedang libur
	fmt.Println("Doctor is on holiday.")
	return true, nil
}

func DoctorActiveSchedule(doctorID uuid.UUID, date string, session string) (model.Consultation, error) {
	var doctorHoliday model.Consultation

	// Cari jadwal dokter pada tanggal dan sesi tertentu
	tx := database.DB.Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session)
	if tx.Error != nil {
		return doctorHoliday, tx.Error
	}

	// Ubah status doctor_available menjadi false
	tx = database.DB.Model(&doctorHoliday).Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).Update("doctor_available", true).Find(&doctorHoliday)
	if tx.Error != nil {
		return doctorHoliday, tx.Error
	}

	// Delete Record sebelumnya
	// tx = database.DB.Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).Delete(&model.Consultation{})
	// if tx.Error != nil {
	// 	return doctorHoliday, tx.Error
	// }

	// Delete DoctorHoliday record
	err := DeleteDoctorHoliday(doctorID)
	if err != nil {
		return doctorHoliday, err
	}

	return doctorHoliday, nil
}

func DeleteDoctorHoliday(doctorID uuid.UUID) error {
	var doctorHoliday model.DoctorHoliday
	if err := database.DB.Where("doctor_id = ?", doctorID).Unscoped().Delete(&doctorHoliday).Error; err != nil {
		return err
	}

	return nil
}
