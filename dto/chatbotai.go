package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

// User
type HealthRecommendationRequest struct {
	PatientID uuid.UUID `json:"patient_id"`
	Message   string    `json:"message"`
}

// Doctor
type HealthRecommendationDoctorRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	SessionID uuid.UUID `json:"session_id"`
	Message   string    `json:"message"`
}

// User
type HealthRecommendationResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

// Doctor
type DoctorHealthRecommendationResponse struct {
	SessionID uuid.UUID `json:"session_id"`
	Status    string    `json:"status"`
	Data      string    `json:"data"`
}

// User
type HealthRecommendationHistoryResponse struct {
	ID        uuid.UUID `json:"id"`
	PatientID uuid.UUID `json:"patient_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
}

// Doctor
type HealthRecommendationHistoryDoctorResponse struct {
	ID        uuid.UUID `json:"id"`
	DoctorID  uuid.UUID `json:"doctor_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	SessionID uuid.UUID `json:"session_id"`
}

// User
func ConvertToHealthRecommendationHistoryResponse(healthRecommendation model.HealthRecommendation) HealthRecommendationHistoryResponse {
	return HealthRecommendationHistoryResponse{
		ID:        healthRecommendation.ID,
		PatientID: healthRecommendation.PatientID,
		Question:  healthRecommendation.Question,
		Answer:    healthRecommendation.Answer,
	}
}

// Doctor
func ConvertToHealthRecommendationHistoryDoctorResponse(doctorHealthRecommendation model.DoctorHealthRecommendation) HealthRecommendationHistoryDoctorResponse {
	return HealthRecommendationHistoryDoctorResponse{
		ID:        doctorHealthRecommendation.ID,
		DoctorID:  doctorHealthRecommendation.DoctorID,
		Question:  doctorHealthRecommendation.Question,
		Answer:    doctorHealthRecommendation.Answer,
		SessionID: doctorHealthRecommendation.SessionID,
	}
}
