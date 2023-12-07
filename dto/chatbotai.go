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
	Status    string    `json:"status"`
	Data      struct {
		ID        uuid.UUID `json:"id"`
		TitleChat string    `json:"titleChat"`
		Tgl       string    `json:"tgl"`
		Pesan     []struct {
			ID       uuid.UUID `json:"id"`
			Pesan    string    `json:"pesan"`
			Waktu    string    `json:"waktu"`
			Pengirim string    `json:"pengirim"`
		} `json:"pesan"`
	} `json:"data"`
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
	response := HealthRecommendationHistoryDoctorResponse{
		Status:    "success",
		Data: struct {
			ID        uuid.UUID `json:"id"`
			TitleChat string    `json:"titleChat"`
			Tgl       string    `json:"tgl"`
			Pesan     []struct {
				ID       uuid.UUID `json:"id"`
				Pesan    string    `json:"pesan"`
				Waktu    string    `json:"waktu"`
				Pengirim string    `json:"pengirim"`
			} `json:"pesan"`
		}{
			ID:        doctorHealthRecommendation.SessionID,
			TitleChat: "",
			Tgl:       doctorHealthRecommendation.CreatedAt.Format("02/01/2006"),
			Pesan: []struct {
				ID       uuid.UUID `json:"id"`
				Pesan    string    `json:"pesan"`
				Waktu    string    `json:"waktu"`
				Pengirim string    `json:"pengirim"`
			}{
				{
					ID:       doctorHealthRecommendation.ID,
					Pesan:    doctorHealthRecommendation.Answer,
					Waktu:    doctorHealthRecommendation.CreatedAt.Format("02/01/2006"),
					Pengirim: "",
				},
			},
		},
	}

	return response
}
