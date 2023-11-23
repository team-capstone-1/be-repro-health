package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetDoctorProfile(id uuid.UUID) (model.Doctor, error) {
	var profileDoctor model.Doctor

	tx := database.DB.Preload("Specialist").Preload("Clinic").Where("id = ?", id).First(&profileDoctor)

	if tx.Error != nil {
		return profileDoctor, tx.Error
	}

	return profileDoctor, nil
}

// Work History

func GetDoctorWorkHistory(id uuid.UUID) ([]model.DoctorWorkHistory, error) {
	var workHistory []model.DoctorWorkHistory

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&workHistory)

	if tx.Error != nil {
		return workHistory, tx.Error
	}

	return workHistory, nil
}

func GetDoctorWorkHistoryByID(id uuid.UUID) (model.DoctorWorkHistory, error) {
	var workHistory model.DoctorWorkHistory

	tx := database.DB.First(&workHistory, id)
	if tx.Error != nil {
		return workHistory, tx.Error
	}

	return workHistory, nil
}

func InsertDoctorWorkHistory(data model.DoctorWorkHistory) (model.DoctorWorkHistory, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.DoctorWorkHistory{}, tx.Error
	}
	return data, nil
}

func DeleteDoctorWorkHistoryByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.DoctorWorkHistory{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateDoctorWorkHistoryByID(id uuid.UUID, updateData model.DoctorWorkHistory) (model.DoctorWorkHistory, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.DoctorWorkHistory{}, tx.Error
	}
	return updateData, nil
}

// Education

func GetDoctorEducation(id uuid.UUID) ([]model.DoctorEducation, error) {
	var education []model.DoctorEducation

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&education)

	if tx.Error != nil {
		return education, tx.Error
	}

	return education, nil
}

func GetDoctorEducationByID(id uuid.UUID) (model.DoctorEducation, error) {
	var education model.DoctorEducation

	tx := database.DB.First(&education, id)
	if tx.Error != nil {
		return education, tx.Error
	}

	return education, nil
}

func InsertDoctorEducation(data model.DoctorEducation) (model.DoctorEducation, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.DoctorEducation{}, tx.Error
	}
	return data, nil
}

func UpdateDoctorEducationByID(id uuid.UUID, updateData model.DoctorEducation) (model.DoctorEducation, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.DoctorEducation{}, tx.Error
	}
	return updateData, nil
}

func DeleteDoctorEducationByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.DoctorEducation{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Certification

func GetDoctorCertification(id uuid.UUID) ([]model.DoctorCertification, error) {
	var certification []model.DoctorCertification

	tx := database.DB.Where("doctor_profile_id = ?", id).Find(&certification)

	if tx.Error != nil {
		return certification, tx.Error
	}

	return certification, nil
}
