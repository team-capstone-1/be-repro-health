package controller

import (
	"net/http"

	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"capstone-project/repository"
	"capstone-project/constant"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateConsultationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	consultation := dto.ConsultationRequest{}
	errBind := c.Bind(&consultation)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(consultation.PatientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse": err.Error(),
		})
	}
	if checkPatient.UserID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	consultationData := dto.ConvertToConsultationModel(consultation)

	clinicData, err := repository.GetClinicByDoctorID(consultationData.DoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get clinic",
			"response": err.Error(),
		})
	}

	consultationData.ClinicID = clinicData.ID

	responseData, err := repository.InsertConsultation(consultationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create consultation",
			"response": err.Error(),
		})
	}

	responseData, _ = repository.GetConsultationByID(responseData.ID)

	consultationResponse := dto.ConvertToConsultationResponse(responseData)

	err = generateTransaction(consultationResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed create transaction",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create new consultation",
		"response": consultationResponse,
	})
}

func generateTransaction(consultation dto.ConsultationResponse) error {
	invoice, date, err := repository.GenerateNextInvoice()
	if err != nil {
		return err
	}

	transaction := model.Transaction{
		ID: uuid.New(),
		ConsultationID: consultation.ID,
		Date: date,
		Invoice: invoice,
		Price: consultation.Doctor.Price,
		AdminPrice: constant.ADMIN_FEE,
		Total: consultation.Doctor.Price + constant.ADMIN_FEE,
		Status: model.Processed,
		PaymentStatus: "pending",
	}

	_, err = repository.InsertTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}