package dto

import (
	"time"
	
	"capstone-project/model"

	"github.com/google/uuid"
)

type PatientRequest struct {
	Name               string    `json:"name" form:"name"`
	TelephoneNumber    string    `json:"telephone_number" form:"telephone_number"`
	ProfileImage       string    `json:"profile_image" form:"profile_image"`
	DateOfBirth        time.Time `json:"date_of_birth" form:"date_of_birth"`
	Relation           string    `json:"relation" form:"relation"`
	Weight             float64   `json:"weight" form:"weight"`
	Height             float64   `json:"height" form:"height"`
	Gender             string    `json:"gender" form:"gender"`
}

type PatientResponse struct {
	ID                 uuid.UUID `json:"id"`
	UserID             uuid.UUID `json:"user_id"`
	Name               string    `json:"name"`
	TelephoneNumber    string    `json:"telephone_number"`
	ProfileImage       string    `json:"profile_image"`
	DateOfBirth        time.Time `json:"date_of_birth"`
	Relation           string    `json:"relation"`
	Weight             float64   `json:"weight"`
	Height             float64   `json:"height"`
	Gender             string    `json:"gender"`
	CreatedAt          time.Time `json:"created_at"`
}

type PatientDashboardResponse struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
}

func ConvertToPatientModel(patient PatientRequest) model.Patient {
	return model.Patient{
		ID:                 uuid.New(),
		Name:               patient.Name,
		TelephoneNumber:    patient.TelephoneNumber,
		ProfileImage:       patient.ProfileImage,
		DateOfBirth:        patient.DateOfBirth,
		Relation:           patient.Relation,
		Weight:             patient.Weight,
		Height:             patient.Height,
		Gender:             patient.Gender,
	}
}

func ConvertToPatientResponse(patient model.Patient) PatientResponse {
	return PatientResponse{
		ID:                 patient.ID,
		UserID:             patient.UserID,
		Name:               patient.Name,
		TelephoneNumber:    patient.TelephoneNumber,
		ProfileImage:       patient.ProfileImage,
		DateOfBirth:        patient.DateOfBirth,
		Relation:           patient.Relation,
		Weight:             patient.Weight,
		Height:             patient.Height,
		Gender:             patient.Gender,
		CreatedAt:          patient.CreatedAt,
	}
}

func ConvertToPatientDashboardResponse(patient model.Patient) PatientResponse {
	return PatientResponse{
		ID:           patient.ID,
		UserID:       patient.UserID,
	}
}
