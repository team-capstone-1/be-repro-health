package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
)

func GetClinicsController(c echo.Context) error {
	responseData, err := repository.GetAllClinics()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get clinics",
			"response":   err.Error(),
		})
	}

	var clinicResponse []dto.ClinicResponse
	for _, clinic := range responseData {
		clinicResponse = append(clinicResponse, dto.ConvertToClinicResponse(clinic))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get clinics",
		"response":   clinicResponse,
	})
}