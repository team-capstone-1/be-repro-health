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

// Doctor
type DoctorHealthRecommendationResponse struct {
	SessionID uuid.UUID `json:"session_id"`
	Status    string    `json:"status"`
	Data      string    `json:"data"`
}

// User
type HealthRecommendationHistoryResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Question      string    `json:"question"`
	Answer        string    `json:"answer"`
	UserSessionID uuid.UUID `json:"user_session_id"`
	IsAIAssistant bool      `json:"is_ai_assistant"`
}

// General
type HealthRecommendationMessage struct {
	ID       uuid.UUID `json:"id"`
	Pesan    string    `json:"pesan"`
	Jawaban  string    `json:"jawaban"`
	Waktu    string    `json:"waktu"`
	Pengirim string    `json:"pengirim"`
}

// Doctor
type HealthRecommendationHistoryDoctorResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID        uuid.UUID                     `json:"id"`
		TitleChat string                        `json:"titleChat"`
		Tgl       string                        `json:"tgl"`
		Pesan     []HealthRecommendationMessage `json:"pesan"`
	} `json:"data"`
}

// User
type HealthRecommendationHistoryUserResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID        uuid.UUID                     `json:"id"`
		TitleChat string                        `json:"titleChat"`
		Tgl       string                        `json:"tgl"`
		Pesan     []HealthRecommendationMessage `json:"pesan"`
	} `json:"data"`
}

// User
func ConvertToHealthRecommendationHistoryUserResponse(healthRecommendations []model.HealthRecommendation) []HealthRecommendationHistoryUserResponse {
	var userresponse []HealthRecommendationHistoryUserResponse

	userMessageMap := make(map[uuid.UUID][]HealthRecommendationMessage)

	for _, userrecommendation := range healthRecommendations {
		userSessionID := userrecommendation.UserSessionID
		userID := userrecommendation.UserID
		userName, err := getUserName(userID)
		if err != nil {
			continue
		}

		message := HealthRecommendationMessage{
			ID:       userrecommendation.ID,
			Pesan:    userrecommendation.Question,
			Jawaban:  userrecommendation.Answer,
			Waktu:    userrecommendation.CreatedAt.Format("02/01/2006 15:04:05"),
			Pengirim: userName,
		}

		userMessageMap[userSessionID] = append(userMessageMap[userSessionID], message)
	}

	// Iterate over the userMessageMap and create HealthRecommendationHistoryUserResponse
	for userSessionID, messages := range userMessageMap {
		userresponse = append(userresponse, HealthRecommendationHistoryUserResponse{
			Status: "success",
			Data: struct {
				ID        uuid.UUID                     `json:"id"`
				TitleChat string                        `json:"titleChat"`
				Tgl       string                        `json:"tgl"`
				Pesan     []HealthRecommendationMessage `json:"pesan"`
			}{
				ID:        userSessionID,
				TitleChat: getChatTitle(messages[0].Pesan),
				Tgl:       messages[0].Waktu,
				Pesan:     messages,
			},
		})
	}

	return userresponse
}

// Doctor
func ConvertToHealthRecommendationHistoryDoctorResponse(doctorHealthRecommendations []model.DoctorHealthRecommendation) []HealthRecommendationHistoryDoctorResponse {
	var response []HealthRecommendationHistoryDoctorResponse

	messageMap := make(map[uuid.UUID][]HealthRecommendationMessage)

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
				TitleChat string                        `json:"titleChat"`
				Tgl       string                        `json:"tgl"`
				Pesan     []HealthRecommendationMessage `json:"pesan"`
			}{
				ID:        sessionID,
				TitleChat: getChatTitle(messages[0].Pesan),
				Tgl:       messages[0].Waktu,
				Pesan:     messages,
			},
		})
	}

	return response
}

// General
func getChatTitle(question string) string {
	words := strings.Fields(question)
	if len(words) > 0 && len(words) >= 3 {
		return strings.Join(words[:3], " ")
	} else if len(words) > 0 && len(words) < 3 {
		return strings.Join(words[:len(words)], " ")
	}
	return "Default Title"
}

// Doctor
func getDoctorName(doctorID uuid.UUID) (string, error) {
	doctor := getDoctorFromDatabase(doctorID)
	if doctor != nil {
		return doctor.Name, nil
	}
	return "", errors.New("Doctor not found")
}

// User
func getUserName(userID uuid.UUID) (string, error) {
	user := getUserFromDatabase(userID)
	if user != nil {
		return user.Name, nil
	}
	return "", errors.New("User not found")
}

// User
func getUserID(userID uuid.UUID) (uuid.UUID, error) {
	user := getUserFromDatabase(userID)
	if user != nil {
		return user.ID, nil
	}
	return uuid.Nil, errors.New("User not found")
}

// Doctor
func getDoctorFromDatabase(doctorID uuid.UUID) *model.Doctor {
	doctor := repository.GetDoctorByIDForAI(doctorID)

	if doctor != nil {
		return doctor
	}
	return nil
}

// User
func getUserFromDatabase(userID uuid.UUID) *model.User {
	user := repository.GetUserByIDForAI(userID)

	if user != nil {
		return user
	}
	return nil
}
