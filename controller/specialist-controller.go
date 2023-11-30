package controller

import (
	"errors"
	"net/http"

	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetSpecialistsController(c echo.Context) error {
	responseData, err := repository.GetSpecialists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed get specialists",
			"response": err.Error(),
		})
	}

	if len(responseData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "no specialists found",
			"response": nil,
		})
	}

	var specialistResponses []dto.SpecialistResponse
	for _, specialist := range responseData {
		specialistResponses = append(specialistResponses, dto.ConvertToSpecialistResponse(specialist))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get specialists",
		"response": specialistResponses,
	})
}

func CreateSpecialistController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	specialistRequest := dto.SpecialistRequest{}
	errBind := c.Bind(&specialistRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "bad request",
			"response": errBind.Error(),
		})
	}

	if user != user {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admins can create specialist.",
		})
	}

	if err := validateDoctorSpecialistRequest(specialistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "Invalid body",
			"response": err.Error(),
		})
	}

	specialistData := dto.ConvertToSpecialistModel(specialistRequest)

	responseData, err := repository.InsertSpecialist(specialistData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed create specialist",
			"response": err.Error(),
		})
	}

	specialistResponse := dto.ConvertToSpecialistResponse(responseData)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create specialist",
		"response": specialistResponse,
	})
}

func UpdateSpecialistController(c echo.Context) error {
	specialistID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "Invalid specialist ID",
			"response": err.Error(),
		})
	}

	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Admin is not valid.",
		})
	}

	specialistData, err := repository.GetSpecialistByID(specialistID)
    if err != nil {
        if err == errors.New("record not found") {
            return c.JSON(http.StatusNotFound, map[string]any{
                "message": "specialist not found",
                "response": nil,
            })
        }

        return c.JSON(http.StatusInternalServerError, map[string]any{
            "message": "failed get specialist",
            "response": err.Error(),
        })
    }

	specialistRequest := dto.SpecialistRequest{}
	errBind := c.Bind(&specialistRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	if err := validateDoctorSpecialistRequest(specialistRequest); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]any{
            "message": "Invalid body",
            "response": err.Error(),
        })
    }

	specialistData.Name = specialistRequest.Name
	specialistData.Image = specialistRequest.Image

	updateSpecialist, err := repository.UpdateSpecialistDoctorByID(specialistID, specialistData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed update specialist",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update specialist",
		"response": updateSpecialist,
	})
}

func DeleteSpecialistController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	specialistID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "Invalid specialist ID",
			"response": err.Error(),
		})
	}

	specialist, err := repository.GetSpecialistByID(specialistID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]any{
            "message": "failed check specialist",
            "response": err.Error(),
        })
    }

	if specialist.ID != specialist.ID {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to delete other user's specialist.",
		})
	}

	err = repository.DeleteSpecialistByID(specialistID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete specialist",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete specialist",
		"response": "success delete specialist with id " + specialistID.String(),
	})
}

func validateDoctorSpecialistRequest(specialist dto.SpecialistRequest) error {

	if specialist.Name == "" || specialist.Image == "" {
		return errors.New("All fields must be filled in")
	}

	return nil
}
