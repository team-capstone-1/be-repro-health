package dto

// HealthRecommendationRequest represents the request structure for health recommendations.
type HealthRecommendationRequest struct {
	Message string `json:"message"`
}


// HealthRecommendationResponse represents the response structure for health recommendations.
type HealthRecommendationResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
