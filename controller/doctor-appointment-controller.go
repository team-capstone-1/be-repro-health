package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func DoctorGetDetailsTransactionController(c echo.Context) error {
	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.DoctorGetDetailsTransaction(uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get details transaction",
			"response": err.Error(),
		})
	}

	if responseData.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "transaction not found",
			"response": nil,
		})
	}

	detailsTransaction := dto.ConvertToDoctorGetDetailsTransactionResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get details transaction",
		"response": detailsTransaction,
	})
}

func DoctorGetDetailsPatientController(c echo.Context) error {
	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	consultation, err := repository.DoctorGetDetailsConsultation(uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get details consultation",
			"response": err.Error(),
		})
	}

	transactions, err := repository.DoctorGetTransactionsForConsultation(uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get transactions for consultation",
			"response": err.Error(),
		})
	}

	consultation.Transaction = transactions

	consultationResponse := dto.ConvertToDoctorGetDetailsPatientResponse(consultation)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get details consultation",
		"response": consultationResponse,
	})
}

func DoctorGetAllConsultations(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	consultations, err := repository.DoctorGetAllConsultations(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get all consultations",
			"response": err.Error(),
		})
	}

	if len(consultations) == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "no consultations found",
			"response": nil,
		})
	}

	var consultationsResponse []dto.DoctorGetAllConsultations
	for _, consultation := range consultations {
		consultationResponse := dto.ConvertToDoctorGetAllConsultations(consultation)
		consultationsResponse = append(consultationsResponse, consultationResponse)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get all consultations",
		"response": consultationsResponse,
	})
}
