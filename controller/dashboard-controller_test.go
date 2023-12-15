package controller_test

// import (
// 	"capstone-project/controller"
// 	"capstone-project/middleware"
// 	"capstone-project/model"
// 	"capstone-project/repository"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestGetDataCountForDoctorControllerOneMonth(t *testing.T) {
// 	// Create a mock repository or use a testing framework
// 	mockRepo := &repository.DoctorAIRepositoryImpl{} // Replace with your actual mock repository

// 	// Mock middleware
// 	mockMiddleware := &middleware.MiddlewareMock{}

// 	// Create a new Echo instance
// 	e := echo.New()

// 	// Create a request
// 	req := httptest.NewRequest(http.MethodGet, "/api/doctor/dashboard/month", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Set the doctor ID in the context
// 	doctorID := uuid.New()
// 	mockMiddleware.On("ExtractTokenUserId", c).Return(doctorID).Once()

// 	// Set the mock repository in the controller
// 	controller := controller.GetDataCountForDoctorControllerOneMonth
// 	controller. = mockRepo
// 	controller.Middleware = mockMiddleware

// 	// Mock repository responses
// 	mockRepo.On("GetConsultationByDoctorAndMonth", doctorID, mock.AnythingOfType("time.Time")).Return([]model.Consultation{}, nil)
// 	mockRepo.On("GetPatientByDoctorAndMonth", doctorID, mock.AnythingOfType("time.Time")).Return([]model.Patient{}, nil)
// 	mockRepo.On("GetAllTransactionsByDoctorID", doctorID).Return([]model.Transaction{}, nil)
// 	mockRepo.On("DoctorGetAllArticlesByMonth", doctorID, mock.AnythingOfType("time.Time")).Return([]model.Article{}, nil)

// 	// Call the controller function
// 	err := controller.GetDataCountForDoctorControllerOneMonth(c)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	// Add more assertions based on your response payload
// 	// ...

// 	// Verify that the mock repository methods were called
// 	mockRepo.AssertExpectations(t)
// }

// // Similar tests for GetDataCountForDoctorControllerOneWeek and GetDataCountForDoctorControllerOneDay...
