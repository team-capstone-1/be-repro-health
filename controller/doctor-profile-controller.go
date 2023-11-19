package controller

import (
	"capstone-project/dto"
	"capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDoctorProfileController(c echo.Context) error {
	id, err := middleware.ExtractTokenDoctor(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "this route is only for doctor",
			"response": err,
		})
	}

	responseData, err := repository.GetDoctorProfile(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor profile",
			"response": err.Error(),
		})
	}

	doctorResponse := dto.ConvertToDoctorProfileResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor profile",
		"response": doctorResponse,
	})
}
