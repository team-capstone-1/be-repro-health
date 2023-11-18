package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
)

func CreateConsultationController(c echo.Context) error {
	consultation := dto.ConsultationRequest{}
	errBind := c.Bind(&consultation)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	consultationData := dto.ConvertToConsultationModel(consultation)
	
	clinicData, err := repository.GetClinicByDoctorID(consultationData.DoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get clinic",
			"response":  err.Error(),
		})
	}

	consultationData.ClinicID = clinicData.ID

	responseData, err := repository.InsertConsultation(consultationData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create consultation",
			"response":  err.Error(),
		})
	}

	consultationResponse := dto.ConvertToConsultationResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new consultation",
		"response":    consultationResponse,
	})
}