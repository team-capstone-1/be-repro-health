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
	consultationData.QueueNumber, err = repository.GenerateQueueNumber(consultationData.Date, consultationData.Session)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed generate queue number",
			"response": err.Error(),
		})
	}

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

	consultationResponse := dto.ConvertToUserConsultationResponse(responseData)

	transaction,err := generateTransaction(consultationResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed create transaction",
			"response": err.Error(),
		})
	}
	consultationResponse.TransactionID = transaction.ID

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create new consultation",
		"response": consultationResponse,
	})
}

func generateTransaction(consultation dto.UserConsultationResponse) (model.Transaction, error) {
	invoice, err := repository.GenerateNextInvoice()
	if err != nil {
		return model.Transaction{},err
	}

	var payment_method = "done"
	if consultation.PaymentMethod != "clinic_payment"{
		payment_method = "pending"
	}

	transaction := model.Transaction{
		ID: uuid.New(),
		ConsultationID: consultation.ID,
		Invoice: invoice,
		Price: consultation.Doctor.Price,
		AdminPrice: constant.ADMIN_FEE,
		Total: consultation.Doctor.Price + constant.ADMIN_FEE,
		Status: model.Waiting,
		PaymentStatus: payment_method,
	}

	transaction, err = repository.InsertTransaction(transaction)
	if err != nil {
		return model.Transaction{},err
	}
	return transaction, nil
}