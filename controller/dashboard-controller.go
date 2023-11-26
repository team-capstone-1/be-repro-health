package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetConsultationSchedulesForDoctorDashboardController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetConsultationsByDoctorIDDahsboard(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get consultations",
			"response": err.Error(),
		})
	}

	var consultationResponse []dto.ConsultationResponse
	for _, doctor := range responseData {
		consultationResponse = append(consultationResponse, dto.ConvertToConsultationResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get consultations",
		"response": consultationResponse,
	})
}

func GetPatientsForDoctorDashboardController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetAllPatientsDashboard(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get patients",
			"response": err.Error(),
		})
	}

	var PatientDashboardResponse []dto.PatientDashboardResponse
	for _, patient := range responseData {
		PatientDashboardResponse = append(PatientDashboardResponse, dto.ConvertToPatientDashboardResponse(patient))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get patients",
		"response": PatientDashboardResponse,
	})
}

func GetTransactionsForDoctorDashboardController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetAllTransactions(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get transaction data",
			"response": err.Error(),
		})
	}

	var transactionResponse []dto.TransactionResponse
	for _, transaction := range responseData {
		transactionResponse = append(transactionResponse, dto.ConvertToTransactionDashboardResponse(transaction))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transaction data",
		"response": transactionResponse,
	})
}

func GetArticleForDoctorDashboardController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetAllArticleDashboard(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get article data",
			"response": err.Error(),
		})
	}

	var articleDashboardResponse []dto.DoctorArticleResponse
	for _, article := range responseData {
		articleDashboardResponse = append(articleDashboardResponse, dto.ConvertToDoctorArticleDashboardResponse(article))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transaction data",
		"response": articleDashboardResponse,
	})
}
