package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"time"

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

func DoctorInactiveSchedule(doctorID uuid.UUID, data model.DoctorHoliday) (model.DoctorHoliday, error) {
	tx := database.DB.Where("doctor_id = ?", doctorID).Save(&data)
	if tx.Error != nil {
		return model.DoctorHoliday{}, tx.Error
	}
	return data, nil
}

func GetDoctorHolidaysByDateAndSession(doctorID uuid.UUID, date time.Time, session string) ([]model.DoctorHoliday, error) {
	var holidays []model.DoctorHoliday
	err := database.DB.Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).Find(&holidays).Error
	return holidays, err
}

func GetConsultationsByDoctorSchedule(doctorID uuid.UUID, date time.Time, session string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	err := database.DB.Where("doctor_id = ? AND date = ? AND session = ?", doctorID, date, session).Find(&consultations).Error
	return consultations, err
}
