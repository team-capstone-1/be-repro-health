package util

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse adalah struktur untuk representasi respons kesalahan.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse membuat instansiasi ErrorResponse baru dengan pesan yang diberikan.
func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

// RespondWithError mengirim respons kesalahan dalam format JSON.
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := NewErrorResponse(message)
	json.NewEncoder(w).Encode(errorResponse)
}
