package controller_test

import (
	"capstone-project/controller"
	"capstone-project/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewDoctorAIController(t *testing.T) {
	mockRepo := &repository.DoctorAIRepositoryImpl{}
	doctorAIController := controller.NewDoctorAIController(mockRepo)

	assert.NotNil(t, doctorAIController)
	assert.Equal(t, mockRepo, doctorAIController.DoctorAIRepo)
}

func TestDoctorAIController_GetHealthRecommendationDoctorHistory(t *testing.T) {
	// Positive Test Case
	// Create a mock repository or use a testing framework
	mockRepo := &repository.DoctorAIRepositoryImpl{}
	controller := controller.NewDoctorAIController(mockRepo)

	e := echo.New()

	// Case: Valid Doctor ID
	req := httptest.NewRequest(http.MethodGet, "/api/doctor/history/123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/doctor/history/:doctor_id")
	c.SetParamNames("doctor_id")
	c.SetParamValues("123")

	err := controller.GetHealthRecommendationDoctorHistory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Negative Test Cases
	// Case: Invalid Doctor ID (non-numeric)
	// req = httptest.NewRequest(http.MethodGet, "/api/doctor/history/abc", nil)
	// rec = httptest.NewRecorder()
	// c = e.NewContext(req, rec)
	// c.SetPath("/api/doctor/history/:doctor_id")
	// c.SetParamNames("doctor_id")
	// c.SetParamValues("abc")

	// err = controller.GetHealthRecommendationDoctorHistory(c)
	// assert.Error(t, err)
	// assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Case: Missing Doctor ID
	// req = httptest.NewRequest(http.MethodGet, "/api/doctor/history/", nil)
	// rec = httptest.NewRecorder()
	// c = e.NewContext(req, rec)
	// c.SetPath("/api/doctor/history/:doctor_id")

	// err = controller.GetHealthRecommendationDoctorHistory(c)
	// assert.Error(t, err)
	// assert.Equal(t, http.StatusBadRequest, rec.Code)

	// // Case: Nonexistent Doctor ID
	// req = httptest.NewRequest(http.MethodGet, "/api/doctor/history/999", nil)
	// rec = httptest.NewRecorder()
	// c = e.NewContext(req, rec)
	// c.SetPath("/api/doctor/history/:doctor_id")
	// c.SetParamNames("doctor_id")
	// c.SetParamValues("999")

	// err = controller.GetHealthRecommendationDoctorHistory(c)
	// assert.Error(t, err)
	// assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDoctorAIController_GetHealthRecommendationDoctorHistoryFromSession(t *testing.T) {
	mockRepo := &repository.DoctorAIRepositoryImpl{} // Create a mock repository or use a testing framework
	controller := controller.NewDoctorAIController(mockRepo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/doctor/history/session/123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/doctor/history/session/:session_id")
	c.SetParamNames("session_id")
	c.SetParamValues("123")

	err := controller.GetHealthRecommendationDoctorHistoryFromSession(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
