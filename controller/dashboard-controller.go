package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"capstone-project/repository"
	"net/http"
	"time"

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

	responseData, err := repository.GetConsultationsByDoctorID(doctor)
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

func GetConsultationCountDoctor(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetConsultationsByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get consultations",
			"response": err.Error(),
		})
	}

	totalConsultation := len(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get consultations",
		"total":   totalConsultation,
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

	responseData, err := repository.GetPatientByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get patients data",
			"response": err.Error(),
		})
	}

	var patientResponse []dto.PatientResponse
	for _, doctor := range responseData {
		patientResponse = append(patientResponse, dto.ConvertToPatientDashboardResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get patients data",
		"response": patientResponse,
	})
}

func GetPatientsForDoctorCountController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetPatientByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get patients data",
			"response": err.Error(),
		})
	}

	totalPatient := len(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get patients data",
		"total":   totalPatient,
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

	responseData, err := repository.GetAllTransactionsByDoctorID(doctor)
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

func GetTransactionsForDoctorCountController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	lastMonth := time.Now().AddDate(0, -2, 0)
	lastMonthData, err := repository.GetDoneTransactionsByDoctorAndMonth(doctor, lastMonth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "failed get last month transaction data",
			"response": err.Error(),
		})
	}

	thisMonth := time.Now().AddDate(0, -1, 0)
	thisMonthData, err := repository.GetDoneTransactionsByDoctorAndMonth(doctor, thisMonth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "failed get this month transaction data",
			"response": err.Error(),
		})
	}

	allTransactions := append(lastMonthData, thisMonthData...)

	totalPriceLastMonth := calculateTotalPrice(lastMonthData)
	totalPriceThisMonth := calculateTotalPrice(thisMonthData)

	var changePercentage float64
	if totalPriceLastMonth > 0 {
		changePercentage = ((totalPriceThisMonth - totalPriceLastMonth) / totalPriceLastMonth) * 100
	} else if totalPriceThisMonth > 0 {
		changePercentage = 100.0
	} else {
		changePercentage = 0.0
	}

	totalEarnings := calculateTotalPrice(allTransactions)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":             "success get data",
		"changePercentage":    changePercentage,
		"totalPriceThisMonth": totalPriceThisMonth,
		"totalPriceLastMonth": totalPriceLastMonth,
		"totalEarnings":       totalEarnings,
	})
}

func calculateTotalPrice(transactions []model.Transaction) float64 {
	totalPrice := 0.0
	for _, transaction := range transactions {
		totalPrice += transaction.Total
	}
	return totalPrice
}

func GetArticleForDoctorDashboardController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	responseData, err := repository.GetAllArticlesByDoctorID(doctor)
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

func GetArticleForDoctorCountController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	allArticles, err := repository.GetAllArticlesByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get all article data",
			"response": err.Error(),
		})
	}

	publishedArticles, err := repository.GetPublishedArticles(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get published article data",
			"response": err.Error(),
		})
	}

	unpublishedArticles, err := repository.GetUnpublishedArticles(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get unpublished article data",
			"response": err.Error(),
		})
	}

	totalAllArticles := len(allArticles)
	totalPublishedArticles := len(publishedArticles)
	totalUnpublishedArticles := len(unpublishedArticles)

	return c.JSON(http.StatusOK, map[string]any{
		"message":                  "success get article data",
		"totalAllArticles":         totalAllArticles,
		"totalPublishedArticles":   totalPublishedArticles,
		"totalUnpublishedArticles": totalUnpublishedArticles,
	})
}
