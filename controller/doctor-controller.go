package controller

import (
	"capstone-project/dto"
	// "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func DoctorLoginController(c echo.Context) error {
	var loginReq = dto.DoctorLoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	data, token, err := repository.CheckDoctor(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "fail login",
			"response": err.Error(),
		})
	}
	response := map[string]any{
		"doctor_id": data.ID,
		"email":     data.Email,
		"token":     token,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success login doctor account",
		"response": response,
	})
}

func SignUpDoctorController(c echo.Context) error {

	var payloads = dto.DoctorSignUpRequest{}
	errBind := c.Bind(&payloads)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	emailExist := repository.CheckDoctorEmail(payloads.Email)
	if emailExist {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "fail sign up",
			"response": "email already exist",
		})
	}

	signUpData := dto.ConvertToDoctorModel(payloads)

	data, err := repository.CreateDoctor(signUpData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "fail sign up",
			"response": err.Error(),
		})
	}

	response := dto.ConvertToDoctorSignUpResponse(data)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create doctor data",
		"response": response,
	})
}

func GetDoctorsController(c echo.Context) error {
	name := c.FormValue("name")

	responseData, err := repository.GetAllDoctors(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctors",
			"response": err.Error(),
		})
	}

	var doctorResponse []dto.DoctorResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctors",
		"response": doctorResponse,
	})
}

func GetDoctorController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetDoctorByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get doctor",
			"reponse": err.Error(),
		})
	}

	doctorResponse := dto.ConvertToDoctorResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctor",
		"response": doctorResponse,
	})
}

func GetDoctorsBySpecialistController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetDoctorsBySpecialist(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctors",
			"response": err.Error(),
		})
	}

	var doctorResponse []dto.DoctorResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctors",
		"response": doctorResponse,
	})
}

func GetDoctorsByClinicController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetDoctorsByClinic(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctors",
			"response": err.Error(),
		})
	}

	var doctorResponse []dto.DoctorResponse
	for _, doctor := range responseData {
		doctorResponse = append(doctorResponse, dto.ConvertToDoctorResponse(doctor))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get doctors",
		"response": doctorResponse,
	})
}
