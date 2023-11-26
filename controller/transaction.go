package controller

import (
	"net/http"
	"fmt"

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
	fmt.Println("tesss")
	uuid, err := uuid.Parse(c.Param("id"))
	fmt.Println(uuid)
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