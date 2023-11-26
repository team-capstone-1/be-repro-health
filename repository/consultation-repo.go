package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetConsultationsByDoctorID(doctorID uuid.UUID) ([]model.Consultation, error) {
	var dataConsultations []model.Consultation

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&dataConsultations)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataConsultations, nil
}


func InsertConsultation(data model.Consultation) (model.Consultation, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}
	
	return data, nil
}

func GetConsultationByID(id uuid.UUID) (model.Consultation, error) {
	var dataconsultation model.Consultation

	tx := database.DB.Preload("Clinic").Preload("Doctor").First(&dataconsultation, id)
	if tx.Error != nil {
		return model.Consultation{}, tx.Error
	}
	return dataconsultation, nil
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