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
// type HealthRecommendationHistoryDoctorResponse struct {
// 	Status    string    `json:"status"`
// 	Data      struct {
// 		ID        uuid.UUID `json:"id"`
// 		TitleChat string    `json:"titleChat"`
// 		Tgl       string    `json:"tgl"`
// 		Pesan     []struct {
// 			ID       uuid.UUID `json:"id"`
// 			Pesan    string    `json:"pesan"`
// 			Waktu    string    `json:"waktu"`
// 			Pengirim string    `json:"pengirim"`
// 		} `json:"pesan"`
// 	} `json:"data"`
// }

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
// func ConvertToHealthRecommendationHistoryDoctorResponse(doctorHealthRecommendation model.DoctorHealthRecommendation) HealthRecommendationHistoryDoctorResponse {
// 	response := HealthRecommendationHistoryDoctorResponse{
// 		Status:    "success",
// 		Data: struct {
// 			ID        uuid.UUID `json:"id"`
// 			TitleChat string    `json:"titleChat"`
// 			Tgl       string    `json:"tgl"`
// 			Pesan     []struct {
// 				ID       uuid.UUID `json:"id"`
// 				Pesan    string    `json:"pesan"`
// 				Waktu    string    `json:"waktu"`
// 				Pengirim string    `json:"pengirim"`
// 			} `json:"pesan"`
// 		}{
// 			ID:        doctorHealthRecommendation.SessionID,
// 			TitleChat: "",
// 			Tgl:       doctorHealthRecommendation.CreatedAt.Format("02/01/2006"),
// 			Pesan: []struct {
// 				ID       uuid.UUID `json:"id"`
// 				Pesan    string    `json:"pesan"`
// 				Waktu    string    `json:"waktu"`
// 				Pengirim string    `json:"pengirim"`
// 			}{
// 				{
// 					ID:       doctorHealthRecommendation.ID,
// 					Pesan:    doctorHealthRecommendation.Answer,
// 					Waktu:    doctorHealthRecommendation.CreatedAt.Format("02/01/2006"),
// 					Pengirim: "",
// 				},
// 			},
// 		},
// 	}

// 	return response
// }

type HealthRecommendationMessage struct {
	ID       uuid.UUID `json:"id"`
	Pesan    string    `json:"pesan"`
	Waktu    string    `json:"waktu"`
	Pengirim string    `json:"pengirim"`
}

type HealthRecommendationHistoryDoctorResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID        uuid.UUID                     `json:"id"`
		TitleChat string                        `json:"titleChat"`
		Tgl       string                        `json:"tgl"`
		Pesan     []HealthRecommendationMessage `json:"pesan"`
	} `json:"data"`
}

func ConvertToHealthRecommendationHistoryDoctorResponse(doctorHealthRecommendations []model.DoctorHealthRecommendation) []HealthRecommendationHistoryDoctorResponse {
	// Create the final response slice
	var response []HealthRecommendationHistoryDoctorResponse

	// Create a map to group messages by session ID
	messageMap := make(map[uuid.UUID][]HealthRecommendationMessage)

	// Iterate over doctorHealthRecommendations and populate messageMap
	for _, recommendation := range doctorHealthRecommendations {
		sessionID := recommendation.SessionID
		message := HealthRecommendationMessage{
			ID:       recommendation.ID,
			Pesan:    recommendation.Answer,
			Waktu:    recommendation.CreatedAt.Format("02/01/2006"),
			Pengirim: "", // Set appropriate value for Pengirim
		}

		messageMap[sessionID] = append(messageMap[sessionID], message)
	}

	// Iterate over the messageMap and create HealthRecommendationHistoryDoctorResponse
	for sessionID, messages := range messageMap {
		response = append(response, HealthRecommendationHistoryDoctorResponse{
			Status: "success",
			Data: struct {
				ID        uuid.UUID                     `json:"id"`
				TitleChat string                        `json:"titleChat"`
				Tgl       string                        `json:"tgl"`
				Pesan     []HealthRecommendationMessage `json:"pesan"`
			}{
				ID:        sessionID,
				TitleChat: "", // Set appropriate value for TitleChat
				Tgl:       messages[0].Waktu,
				Pesan:     messages,
			},
		})
	}

	return response
}
