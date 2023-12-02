package controller

import (
	"capstone-project/dto"
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
		return c.JSON(http.StatusNotFound, map[string]interface{}{
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
