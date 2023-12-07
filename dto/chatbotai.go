package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type HealthRecommendationRequest struct {
	PatientID        uuid.UUID `json:"patient_id"`
	PatientSessionID uuid.UUID `json:"session_id"`
	Message          string    `json:"message"`
}

type HealthRecommendationDoctorRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	SessionID uuid.UUID `json:"session_id"`
	Message   string    `json:"message"`
}

type HealthRecommendationResponse struct {
	PatientSessionID uuid.UUID `json:"session_id"`
	Status           string    `json:"status"`
	Data             string    `json:"data"`
}

type DoctorHealthRecommendationResponse struct {
	SessionID uuid.UUID `json:"session_id"`
	Status    string    `json:"status"`
	Data      string    `json:"data"`
}

type HealthRecommendationHistoryResponse struct {
	ID               uuid.UUID `json:"id"`
	PatientID        uuid.UUID `json:"patient_id"`
	Question         string    `json:"question"`
	Answer           string    `json:"answer"`
	PatientSessionID uuid.UUID `json:"session_id"`
}

type HealthRecommendationHistoryDoctorResponse struct {
	ID        uuid.UUID `json:"id"`
	DoctorID  uuid.UUID `json:"doctor_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	SessionID uuid.UUID `json:"session_id"`
}

func ConvertToHealthRecommendationHistoryResponse(healthRecommendation model.HealthRecommendation) HealthRecommendationHistoryResponse {
	return HealthRecommendationHistoryResponse{
		ID:               healthRecommendation.ID,
		PatientID:        healthRecommendation.PatientID,
		Question:         healthRecommendation.Question,
		Answer:           healthRecommendation.Answer,
		PatientSessionID: healthRecommendation.PatientSessionID,
	}
}

func ConvertToHealthRecommendationHistoryDoctorResponse(doctorHealthRecommendation model.DoctorHealthRecommendation) HealthRecommendationHistoryDoctorResponse {
	return HealthRecommendationHistoryDoctorResponse{
		ID:        doctorHealthRecommendation.ID,
		DoctorID:  doctorHealthRecommendation.DoctorID,
		Question:  doctorHealthRecommendation.Question,
		Answer:    doctorHealthRecommendation.Answer,
		SessionID: doctorHealthRecommendation.SessionID,
	}
}
