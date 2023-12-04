package repository

import (
	"fmt"
	"time"

	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GenerateQueueNumber(date time.Time, session string) (string, error) {
	formattedDate := date.Format("2006-01-02")

	var sequence int
	result := database.DB.Table("consultations").
		Where("date = ? AND session = ?", formattedDate, session).
		Select("COALESCE(MAX(CAST(SUBSTRING(queue_number, 1, 3) AS SIGNED)), 0) AS max_sequence").
		Scan(&sequence)

	if result.Error != nil {
		return "", result.Error
	}

	// Increment the sequence
	sequence++

	formattedQueueNumber := fmt.Sprintf("%03d", sequence)
	fmt.Println(formattedQueueNumber)
	return formattedQueueNumber, nil
}

func GetConsultationsByDoctorID(doctorID uuid.UUID) ([]model.Consultation, error) {
	var dataConsultations []model.Consultation

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&dataConsultations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataConsultations, nil
}

func GetConsultationByID(id uuid.UUID) (model.Consultation, error) {
	var dataconsultation model.Consultation

	tx := database.DB.Preload("Clinic").Preload("Doctor").Preload("Doctor.Specialist").First(&dataconsultation, id)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}
	return dataconsultation, nil
}

func GetConsultationByDoctorAndMonth(doctorID uuid.UUID, month time.Time) ([]model.Consultation, error) {
	var consultations []model.Consultation

	startOfMonth := month.AddDate(0, 0, 1)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	tx := database.DB.
		Preload("Doctor").
		Joins("JOIN transactions ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND consultations.date BETWEEN ? AND ? AND transactions.payment_status = 'done'", doctorID, startOfMonth, endOfMonth).
		Find(&consultations)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return consultations, nil
}

func GetConsultationByDoctorAndWeek(doctorID uuid.UUID, week time.Time) ([]model.Consultation, error) {
	var consultations []model.Consultation

	startOfWeek := week.AddDate(0, 0, 0)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	tx := database.DB.
		Preload("Doctor").
		Joins("JOIN transactions ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND consultations.date BETWEEN ? AND ? AND transactions.payment_status = 'done'", doctorID, startOfWeek, endOfWeek).
		Find(&consultations)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return consultations, nil
}

func InsertConsultation(data model.Consultation) (model.Consultation, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}

	return data, nil
}

func RescheduleConsultation(id uuid.UUID, updateData model.Consultation) (model.Consultation, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}
	return updateData, nil
}

func GetConsultationByTransactionID(transactionID uuid.UUID) (model.Consultation, error) {
	var dataconsultation model.Consultation

	tx := database.DB.Joins("JOIN transactions ON consultations.id = transactions.consultation_id").
		Where("transactions.id = ?", transactionID).
		First(&dataconsultation)

	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}

	return dataconsultation, nil
}
