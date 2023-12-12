package dto

import (
	"capstone-project/model"
	"capstone-project/repository"
	"errors"
	"strings"

	"github.com/google/uuid"
)

// User
type HealthRecommendationRequest struct {
	UserID        uuid.UUID `json:"user_id"`
	UserSessionID uuid.UUID `json:"user_session_id"`
	Message       string    `json:"message"`
}

// Doctor
type HealthRecommendationDoctorRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	SessionID uuid.UUID `json:"session_id"`
	Message   string    `json:"message"`
}

// User
type HealthRecommendationResponse struct {
	UserSessionID uuid.UUID `json:"user_session_id"`
	Status        string    `json:"status"`
	Data          string    `json:"data"`
}

type DoctorHealthRecommendationResponse struct {
	SessionID uuid.UUID `json:"session_id"`
	Status    string    `json:"status"`
	Data      string    `json:"data"`
}

type HealthRecommendationHistoryResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Question      string    `json:"question"`
	Answer        string    `json:"answer"`
	UserSessionID uuid.UUID `json:"user_session_id"`
}

// User
func ConvertToHealthRecommendationHistoryResponse(healthRecommendation model.HealthRecommendation) HealthRecommendationHistoryResponse {
	return HealthRecommendationHistoryResponse{
		ID:            healthRecommendation.ID,
		UserID:        healthRecommendation.UserID,
		Question:      healthRecommendation.Question,
		Answer:        healthRecommendation.Answer,
		UserSessionID: healthRecommendation.UserSessionID,
	}
}

type HealthRecommendationMessage struct {
	ID       uuid.UUID `json:"id"`
	Pesan    string    `json:"pesan"`
	Jawaban  string    `json:"jawaban"`
	Waktu    string    `json:"waktu"`
	Pengirim string    `json:"pengirim"`
}

type HealthRecommendationHistoryDoctorResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID        uuid.UUID                     `json:"id"`
		DoctorID  uuid.UUID                     `json:"doctor_id"`
		TitleChat string                        `json:"titleChat"`
		Tgl       string                        `json:"tgl"`
		Pesan     []HealthRecommendationMessage `json:"pesan"`
	} `json:"data"`
}

type HealthRecommendationHistoryPatientResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID        uuid.UUID                     `json:"id"`
		PatientID uuid.UUID                     `json:"patient_id"`
		TitleChat string                        `json:"titleChat"`
		Tgl       string                        `json:"tgl"`
		Pesan     []HealthRecommendationMessage `json:"pesan"`
	} `json:"data"`
}

// Doctor
func ConvertToHealthRecommendationHistoryDoctorResponse(doctorHealthRecommendations []model.DoctorHealthRecommendation) []HealthRecommendationHistoryDoctorResponse {
	// Create the final response slice
	var response []HealthRecommendationHistoryDoctorResponse

	// Create a map to group messages by session ID
	messageMap := make(map[uuid.UUID][]HealthRecommendationMessage)

	// Iterate over doctorHealthRecommendations and populate messageMap
	for _, recommendation := range doctorHealthRecommendations {
		sessionID := recommendation.SessionID
		doctorID := recommendation.DoctorID
		doctorName, err := getDoctorName(doctorID)
		if err != nil {
			continue
		}

		message := HealthRecommendationMessage{
			ID:       recommendation.ID,
			Pesan:    recommendation.Question,
			Jawaban:  recommendation.Answer,
			Waktu:    recommendation.CreatedAt.Format("02/01/2006 15:04:05"),
			Pengirim: doctorName,
		}

		messageMap[sessionID] = append(messageMap[sessionID], message)
	}

	// Iterate over the messageMap and create HealthRecommendationHistoryDoctorResponse
	for sessionID, messages := range messageMap {
		response = append(response, HealthRecommendationHistoryDoctorResponse{
			Status: "success",
			Data: struct {
				ID        uuid.UUID                     `json:"id"`
				DoctorID  uuid.UUID                     `json:"doctor_id"`
				TitleChat string                        `json:"titleChat"`
				Tgl       string                        `json:"tgl"`
				Pesan     []HealthRecommendationMessage `json:"pesan"`
			}{
				ID:        sessionID,
				DoctorID:  messages[0].ID,
				TitleChat: getChatTitle(messages[0].Pesan),
				Tgl:       messages[0].Waktu,
				Pesan:     messages,
			},
		})
	}

	return response
}

func getChatTitle(question string) string {
	words := strings.Fields(question)
	if len(words) > 0 && len(words) >= 3 {
		return strings.Join(words[:3], " ")
	} else if len(words) > 0 && len(words) < 3{
		return strings.Join(words[:len(words)], " ")
	}
	return "Default Title"
}

func getDoctorName(doctorID uuid.UUID) (string, error) {
	doctor := getDoctorFromDatabase(doctorID)
	if doctor != nil {
		return doctor.Name, nil
	}
	return "", errors.New("Doctor not found")
}

func getDoctorFromDatabase(doctorID uuid.UUID) *model.Doctor {
	doctor := repository.GetDoctorByIDForAI(doctorID)

	if doctor != nil {
		return doctor
	}
	return nil
}

// User
