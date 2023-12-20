package controller_test

import (
	"capstone-project/controller"
	"capstone-project/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDoctorAIController_GetHealthRecommendationDoctorHistory(t *testing.T) {
	e := echo.New()

	doctorAIRepo := &repository.DoctorAIRepositoryImpl{}

	ac := &controller.DoctorAIController{
		DoctorAIRepo: doctorAIRepo,
	}

	doctorID := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/chatbot/health-recommendation/"+doctorID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ac.GetHealthRecommendationDoctorHistory(c)

	c.SetPath("/chatbot/health-recommendation" + doctorID.String())

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDoctorAIController_GetHealthRecommendationDoctorHistoryFromSession(t *testing.T) {
	e := echo.New()

	doctorAIRepo := &repository.DoctorAIRepositoryImpl{}

	ac := &controller.DoctorAIController{
		DoctorAIRepo: doctorAIRepo,
	}

	sessionID := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/chatbot/health-recommendation/session/"+sessionID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ac.GetHealthRecommendationDoctorHistoryFromSession(c)

	c.SetPath("/chatbot/health-recommendation/session" + sessionID.String())

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
