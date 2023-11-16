package controller

import (
	"net/http"
	"strconv"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
)

func GetPatientsController(c echo.Context) error {
	responseData, err := repository.GetAllPatients()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patients",
			"response":   err.Error(),
		})
	}

	var patientResponse []dto.PatientResponse
	for _, patient := range responseData {
		patientResponse = append(patientResponse, dto.ConvertToPatientResponse(patient))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get patients",
		"response":   patientResponse,
	})
}

func GetPatientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	responseData, err := repository.GetPatientByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse":   err.Error(),
		})
	}

	patientResponse := dto.ConvertToPatientResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get patient",
		"response":    patientResponse,
	})
}

func CreatePatientController(c echo.Context) error {
	patient := dto.PatientRequest{}
	errBind := c.Bind(&patient)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	patientData := dto.ConvertToPatientModel(patient)
	
	responseData, err := repository.InsertPatient(patientData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create patient",
			"response":  err.Error(),
		})
	}

	patientResponse := dto.ConvertToPatientResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new patient",
		"response":    patientResponse,
	})
}

func UpdatePatientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	updateData := dto.PatientRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	patientData := dto.ConvertToPatientModel(updateData)

	responseData, err := repository.UpdatePatientByID(id, patientData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed update patient",
			"response":   err.Error(),
		})
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	responseData, err = repository.GetPatientByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse":   err.Error(),
		})
	}

	patientResponse := dto.ConvertToPatientResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update patient",
		"response":    patientResponse,
	})
}

func DeletePatientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	_, err = repository.GetPatientByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete patient",
			"reponse":   err.Error(),
		})
	}

	err = repository.DeletePatientByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed delete patient",
			"reponse":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete patient",
		"response": "success delete patient with id " + strconv.Itoa(id),
	})
}
