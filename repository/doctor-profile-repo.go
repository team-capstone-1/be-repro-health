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

func GetDoctorWorkHistories(id uuid.UUID) ([]model.DoctorWorkHistory, error) {
	var workHistory []model.DoctorWorkHistory

	tx := database.DB.Where("doctor_id = ?", id).Find(&workHistory)

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

func InsertDoctorWorkHistory(workHistory model.DoctorWorkHistory) (model.DoctorWorkHistory, error) {
	tx := database.DB.Create(&workHistory)
	if tx.Error != nil {
		return model.DoctorWorkHistory{}, tx.Error
	}
	return workHistory, nil
}

func DeleteDoctorWorkHistoryByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.DoctorWorkHistory{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateDoctorWorkHistoryByID(id uuid.UUID, updateData model.DoctorWorkHistory) (model.DoctorWorkHistory, error) {
	tx := database.DB.Save(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.DoctorWorkHistory{}, tx.Error
	}
	return updateData, nil
}

// Education

func GetDoctorEducation(id uuid.UUID) ([]model.DoctorEducation, error) {
	var educationHistory []model.DoctorEducation

	tx := database.DB.Where("doctor_id = ?", id).Find(&educationHistory)

	if tx.Error != nil {
		return educationHistory, tx.Error
	}

	return educationHistory, nil
}

func GetDoctorEducationByID(id uuid.UUID) (model.DoctorEducation, error) {
	var educationHistory model.DoctorEducation

	tx := database.DB.First(&educationHistory, id)
	if tx.Error != nil {
		return educationHistory, tx.Error
	}

	return educationHistory, nil
}

func InsertDoctorEducation(education model.DoctorEducation) (model.DoctorEducation, error) {
	tx := database.DB.Create(&education)
	if tx.Error != nil {
		return model.DoctorEducation{}, tx.Error
	}
	return education, nil
}

func DeleteDoctorEducationByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.DoctorEducation{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateDoctorEducationByID(id uuid.UUID, updateData model.DoctorEducation) (model.DoctorEducation, error) {
	tx := database.DB.Save(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.DoctorEducation{}, tx.Error
	}
	return updateData, nil
}

// Certification

func GetDoctorCertification(id uuid.UUID) ([]model.DoctorCertification, error) {
	var certificationHistory []model.DoctorCertification

	tx := database.DB.Where("doctor_id = ?", id).Find(&certificationHistory)

	if tx.Error != nil {
		return certificationHistory, tx.Error
	}

	return certificationHistory, nil
}

func GetDoctorCertificationByID(id uuid.UUID) (model.DoctorCertification, error) {
	var certificationHistory model.DoctorCertification

	tx := database.DB.First(&certificationHistory, id)
	if tx.Error != nil {
		return certificationHistory, tx.Error
	}

	return certificationHistory, nil
}

func InsertDoctorCertification(certification model.DoctorCertification) (model.DoctorCertification, error) {
	tx := database.DB.Create(&certification)
	if tx.Error != nil {
		return model.DoctorCertification{}, tx.Error
	}
	return certification, nil
}

func DeleteDoctorCertificationByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.DoctorCertification{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateDoctorCertificationByID(id uuid.UUID, updateData model.DoctorCertification) (model.DoctorCertification, error) {
	tx := database.DB.Save(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.DoctorCertification{}, tx.Error
	}
	return updateData, nil
}
