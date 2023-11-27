package controller

import (
	"net/http"

	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetSpecialistsController(c echo.Context) error {
	admin := m.ExtractTokenUserId(c)
	if admin == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	responseData, err := repository.GetSpecialists(admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
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
	admin := m.ExtractTokenUserId(c)
	if admin == uuid.Nil {
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
	
	specialistData := dto.ConvertToSpecialistModel(specialistRequest)
	specialistData.ID = admin

	responseData, err := repository.InsertSpecialist(admin, specialistData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create specialist",
			"response": err.Error(),
		})
	}

	specialistResponse := dto.ConvertToSpecialistResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create specialist",
		"response": specialistResponse,
	})
}

func UpdateSpecialistController(c echo.Context) error {
	admin := m.ExtractTokenUserId(c)
	if admin == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "bad request",
			"response": err.Error(),
		})
	}

	checkSpecialist, err := repository.GetSpecialistByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check specialist",
			"response": err.Error(),
		})
	}

	if checkSpecialist.ID != admin {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
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

	specialistData := dto.ConvertToSpecialistModel(specialistRequest)
	specialistData.ID = admin

	responseData, err := repository.UpdateSpecialistDoctorByID(admin, specialistData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update specialist",
			"response": err.Error(),
		})
	}

	responseData, err = repository.GetSpecialistByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get specialist",
			"response": err.Error(),
		})
	}

	specialistResponse := dto.ConvertToSpecialistResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update specialist",
		"response": specialistResponse,
	})
}

func DeleteSpecialistController(c echo.Context) error {
	admin := m.ExtractTokenUserId(c)
	if admin == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "bad request",
			"response": err.Error(),
		})
	}

	checkSpecialist, err := repository.GetSpecialistByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed check specialist",
			"response": err.Error(),
		})
	}

	checkID, err := repository.GetSpecialistByID(checkSpecialist.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete specialist",
			"response": err.Error(),
		})
	}

	if checkID.ID != admin {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	err = repository.DeleteSpecialistByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete specialist",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete specialist",
		"response": nil,
	})
}
