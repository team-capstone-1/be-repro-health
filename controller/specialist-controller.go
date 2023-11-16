package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
)

func GetSpecialistsController(c echo.Context) error {
	responseData, err := repository.GetAllSpecialists()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get specialists",
			"response":   err.Error(),
		})
	}

	var specialistResponse []dto.SpecialistResponse
	for _, specialist := range responseData {
		specialistResponse = append(specialistResponse, dto.ConvertToSpecialistResponse(specialist))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get specialists",
		"response":   specialistResponse,
	})
}