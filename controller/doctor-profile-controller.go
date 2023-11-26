package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetDoctorProfileController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	responseData, err := repository.GetDoctorProfile(user)
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

func GetDoctorWorkHistoriesController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	responseData, err := repository.GetDoctorWorkHistory(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor work histories",
			"response": err.Error(),
		})

	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "no doctor work histories found",
			"response": nil,
		})
	}

	var doctorResponse []dto.DoctorWorkHistoryResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorWorkHistoriesResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor work histories",
		"response": doctorResponse,
	})
}

func GetDoctorEducationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	responseData, err := repository.GetDoctorEducation(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor educations",
			"response": err.Error(),
		})

	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "no doctor education found",
			"response": nil,
		})
	}

	var doctorResponse []dto.DoctorEducationResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorEducationResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor educations",
		"response": doctorResponse,
	})
}

func GetDoctorCertificationController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	responseData, err := repository.GetDoctorCertification(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor certification",
			"response": err.Error(),
		})

	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "no doctor certifications found",
			"response": nil,
		})
	}

	var doctorResponse []dto.DoctorCertificationResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorCertificationResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor certifications",
		"response": doctorResponse,
	})
}
