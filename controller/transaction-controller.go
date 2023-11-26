package controller

import (
	"net/http"
	"errors"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

func GetTransactionController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	transactionResponse := dto.ConvertToTransactionResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transaction",
		"response": transactionResponse,
	})
}

func GetPatientTransactionsController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetPatientTransactions(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get transaction",
			"reponse": err.Error(),
		})
	}

	var transactionResponse []dto.TransactionResponse
	for _, transaction := range responseData {
		transactionResponse = append(transactionResponse, dto.ConvertToTransactionResponse(transaction))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get transaction",
		"response": transactionResponse,
	})
}

func CreatePaymentController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	paymentExist := repository.CheckPayment(uuid)
	if paymentExist{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create payment",
			"response": errors.New("Payment Already Exist").Error(),
		})
	}

	payment := dto.PaymentRequest{}
	errBind := c.Bind(&payment)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	_, err = repository.GetTransactionByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create payment",
			"reponse":   err.Error(),
		})
	}

	paymentData := dto.ConvertToPaymentModel(payment)
	paymentData.TransactionID = uuid
	
	responseData, err := repository.InsertPayment(paymentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create payment",
			"response":  err.Error(),
		})
	}

	paymentResponse := dto.ConvertToPaymentResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new payment",
		"response":    paymentResponse,
	})
}