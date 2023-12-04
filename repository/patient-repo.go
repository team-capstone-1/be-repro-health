package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

func GetAllPatients(user uuid.UUID) ([]model.Patient, error) {
	var datapatients []model.Patient

	tx := database.DB.Where("user_id = ?", user).Find(&datapatients)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return datapatients, nil
}

func GetPatientByDoctorID(doctorID uuid.UUID) ([]model.Patient, error) {
	var datapatients []model.Patient

	tx := database.DB.Joins("JOIN consultations ON patients.id = consultations.patient_id").
		Where("consultations.doctor_id = ?", doctorID).
		Find(&datapatients)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return datapatients, nil
}

func GetPatientByDoctorAndMonth(doctorID uuid.UUID, month time.Time) ([]model.Patient, error) {
	var patients []model.Patient

	startOfMonth := month.AddDate(0, 0, 1)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	tx := database.DB.
		Preload("User").
		Joins("JOIN consultations ON consultations.patient_id = patients.id").
		Joins("JOIN transactions ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND transactions.payment_status = 'done'", doctorID).
		Group("patients.id").
		Having("MIN(consultations.date) BETWEEN ? AND ?", startOfMonth, endOfMonth).
		Find(&patients)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return patients, nil
}

func GetPatientByDoctorAndWeek(doctorID uuid.UUID, week time.Time) ([]model.Patient, error) {
	var patients []model.Patient

	startOfWeek := week.AddDate(0, 0, 0)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	tx := database.DB.
		Preload("User").
		Joins("JOIN consultations ON consultations.patient_id = patients.id").
		Joins("JOIN transactions ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND transactions.payment_status = 'done'", doctorID).
		Group("patients.id").
		Having("MIN(consultations.date) BETWEEN ? AND ?", startOfWeek, endOfWeek).
		Find(&patients)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return patients, nil
}

func GetPatientByID(id uuid.UUID) (model.Patient, error) {
	var datapatient model.Patient

	tx := database.DB.First(&datapatient, id)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return datapatient, nil
}

func InsertPatient(data model.Patient) (model.Patient, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return data, nil
}

func DeletePatientByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Patient{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdatePatientByID(id uuid.UUID, updateData model.Patient) (model.Patient, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.Patient{}, tx.Error
	}
	return updateData, nil
}

func GetProfileByPatientID(id uuid.UUID) string {
	var data model.Patient
	database.DB.Where("id = ?", id).First(&data)

	return data.ProfileImage
}
