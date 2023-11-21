package controller

import (
	"capstone-project/dto"
	"capstone-project/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetConsultationSchedulesForDoctorDashboardController(c echo.Context) error {
	responseData, err := repository.GetAllConsultation()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get consultations data",
			"response": err.Error(),
		})
	}

	var consultationResponse []dto.ConsultationResponse
	for _, consultation := range responseData {
		consultationResponse = append(consultationResponse, dto.ConvertToConsultationResponse(consultation))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get consultations data",
		"response": consultationResponse,
	})
}

func GetPatientsForDoctorDashboardController(c echo.Context) error {
	responseData, err := repository.GetAllPatientsDashboard()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get patients data",
			"response": err.Error(),
		})
	}

	var patientResponse []dto.PatientResponse
	for _, patient := range responseData {
		patientResponse = append(patientResponse, dto.ConvertToPatientDashboardResponse(patient))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get patients data",
		"response": patientResponse,
	})
}

func GetTransactionsForDoctorDashboardController(c echo.Context) error {
	responseData, err := repository.GetAllTransactions()
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
	responseData, err := repository.GetAllArticleDashboard()
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