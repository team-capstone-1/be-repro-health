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

	name := c.FormValue("name")
	status := c.FormValue("status")

	consultations, err := repository.DoctorGetAllConsultations(user, name, status)
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

func DoctorConfirmConsultationController(c echo.Context) error {
	updateData := dto.DoctorConfirmConsultationRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	confirmConsultation := dto.ConvertToDoctorConfirmConsultationModel(updateData, updateData.ConsultationID)

	err := repository.DoctorConfirmConsultation(updateData.ConsultationID, confirmConsultation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed confirm consultation",
			"response": err.Error(),
		})
	}

	consultation, err := repository.DoctorGetDetailsConsultation(updateData.ConsultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get consultation",
			"response": err.Error(),
		})
	}

	transactions, err := repository.DoctorGetTransactionsForConsultation(updateData.ConsultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get transactions for consultation",
			"response": err.Error(),
		})
	}

	consultation.Transaction = transactions

	consultationResponse := dto.ConvertToDoctorGetDetailsPatientResponse(consultation)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success confirm consultation",
		"response": consultationResponse,
	})
}
func DoctorFinishedConsultationController(c echo.Context) error {
	updateData := dto.DoctorConfirmConsultationRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	confirmConsultation := dto.ConvertToDoctorFinishConsultationModel(updateData, updateData.ConsultationID)

	err := repository.DoctorFinishConsultation(updateData.ConsultationID, confirmConsultation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed finish consultation",
			"response": err.Error(),
		})
	}

	consultation, err := repository.DoctorGetDetailsConsultation(updateData.ConsultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get consultation",
			"response": err.Error(),
		})
	}

	transactions, err := repository.DoctorGetTransactionsForConsultation(updateData.ConsultationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get transactions for consultation",
			"response": err.Error(),
		})
	}

	consultation.Transaction = transactions

	consultationResponse := dto.ConvertToDoctorGetDetailsPatientResponse(consultation)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success finish consultation",
		"response": consultationResponse,
	})
}
