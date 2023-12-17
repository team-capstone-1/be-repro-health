package controller

import (
	"capstone-project/config"
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"capstone-project/template"
	"fmt"
	"math/rand"
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

	if loginReq.Email == "" || loginReq.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "email or password cannot be blank",
			"response": nil,
		})
	}

	data, token, err := repository.CheckDoctor(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "doctor account not found",
			"response": err.Error(),
		})
	}

	c.Set("doctor_id", data.ID)

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

func ChangeDoctorPasswordController(c echo.Context) error {
	updatePassword := dto.ChangeDoctorPasswordRequest{}
	errBind := c.Bind(&updatePassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	updatePassword.ID = user
	userData := dto.ConvertToChangeDoctorPasswordModel(updatePassword)

	responseData, err := repository.UpdateDoctorPassword(userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed change password",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success change password",
		"response": "success change password " + responseData.Email,
	})
}

func DoctorSendOTPController(c echo.Context) error {
	var OTPReq = dto.DoctorOTPRequest{}
	errBind := c.Bind(&OTPReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	otp := fmt.Sprint(rand.Intn(99999-10000) + 10000)

	err := repository.SetDoctorOTP(OTPReq.Email, otp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed set otp",
			"response": err.Error(),
		})
	}

	emailBody, err := template.RenderOTPTemplate(otp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed render otp template",
			"response": err.Error(),
		})
	}

	err = config.SendMail(OTPReq.Email, "Reproduction Health Forgot Password OTP", emailBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed send email",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "Success send otp",
		"response": "success send otp to " + OTPReq.Email,
	})
}

func DoctorValidateOTPController(c echo.Context) error {
	var ValidateOTPReq = dto.DoctorValidateOTPRequest{}
	errBind := c.Bind(&ValidateOTPReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	data, token, err := repository.ValidateDoctorOTP(ValidateOTPReq.Email, ValidateOTPReq.OTP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed validate otp",
			"response": err.Error(),
		})
	}
	response := map[string]any{
		"token": token,
		"email": data.Email,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success validate otp",
		"response": response,
	})
}
func GetDoctorsBySpecialistAndClinicController(c echo.Context) error {
	SpecialistID, err := uuid.Parse(c.Param("specialist_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse specialist id",
			"response": err.Error(),
		})
	}
	ClinicID, err := uuid.Parse(c.Param("clinic_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse specialist id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetDoctorsBySpecialistAndClinic(SpecialistID, ClinicID)
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

func SignUpDoctorControllerTesting() echo.HandlerFunc {
	return SignUpDoctorController
}
