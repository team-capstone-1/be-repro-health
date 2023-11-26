package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetConsultationsByDoctorIDDahsboard(doctorID uuid.UUID) ([]model.Consultation, error) {
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