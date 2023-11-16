package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type PatientRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type PatientResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func ConvertToPatientModel(patient PatientRequest) model.Patient {
	return model.Patient{
		Name:     patient.Name,
		Email:    patient.Email,
		Password: patient.Password,
	}
}

func ConvertToPatientResponse(patient model.Patient) PatientResponse {
	return PatientResponse{
		ID:    patient.ID,
		Name:  patient.Name,
		Email: patient.Email,
	}
}
