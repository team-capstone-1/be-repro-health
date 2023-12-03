package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type HealthRecommendationRequest struct {
	PatientID   uuid.UUID `json:"patient_id"`
	Message 	string `json:"message"`
}

type HealthRecommendationResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type HealthRecommendationHistoryResponse struct {
	ID		    uuid.UUID `json:"id"`
	PatientID   uuid.UUID `json:"patient_id"`
	Question 	string `json:"question"`
	Answer 		string `json:"answer"`
}

func ConvertToHealthRecommendationHistoryResponse(healthRecommendation model.HealthRecommendation) HealthRecommendationHistoryResponse {
	return HealthRecommendationHistoryResponse{
		ID:        healthRecommendation.ID,
		PatientID: healthRecommendation.PatientID,
		Question:  healthRecommendation.Question,
		Answer:    healthRecommendation.Answer,
	}
}