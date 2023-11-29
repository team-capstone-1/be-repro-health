package dto

type HealthRecommendationRequest struct {
	Message string `json:"message"`
}

type HealthRecommendationResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
