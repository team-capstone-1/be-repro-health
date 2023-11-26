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
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "bad request",
			"response": err.Error(),
		})
	}

	specialists, err := repository.GetSpecialists(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get specialists",
			"response": err.Error(),
		})
	}

	var specialistsResponse []dto.SpecialistResponse
	for _, specialist := range specialists {
		specialistsResponse = append(specialistsResponse, dto.ConvertToSpecialistResponse(specialist))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get specialists",
		"response": specialists,
	})
}

func CreateSpecialistController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "bad request",
			"response": err.Error(),
		})
	}

	specialist := dto.SpecialistRequest{}
	errBind := c.Bind(&specialist)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	specialistData := dto.ConvertToSpecialistModel(specialist)
	specialistData.ID = uuid

	responseData, err := repository.InsertSpecialist(specialistData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create specialist",
			"response": err.Error(),
		})
	}

	articleResponse := dto.ConvertToSpecialistResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create specialist",
		"response": articleResponse,
	})
}

func UpdateSpecialistController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
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

	if checkSpecialist.ID != doctor {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "user is not valid.",
		})
	}

	specialist := new(dto.SpecialistRequest)
	if err := c.Bind(specialist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "bad request",
			"response": err.Error(),
		})
	}

	specialistModel := dto.ConvertToSpecialistModel(*specialist)
	specialistModel.ID = uuid

	responseData, err := repository.UpdateSpecialist(specialistModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update specialist",
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
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
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

	checkID, err := repository.GetDoctorByID(checkSpecialist.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete specialist",
			"response": err.Error(),
		})
	}

	if checkID.ID != doctor {
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
