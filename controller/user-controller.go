package controller

import (
	"capstone-project/config"
	"capstone-project/dto"
	"capstone-project/repository"
	"capstone-project/template"
	m "capstone-project/middleware"

	"net/http"
	"math/rand"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

func LoginUserController(c echo.Context) error {
	var loginReq = dto.LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	data, token, err := repository.CheckUser(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "fail login",
			"response": err.Error(),
		})
	}
	response := map[string]any{
		"user_id": data.ID,
		"name": data.Name,
		"token":   token,
		"email":   data.Email,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message":  "Success login",
		"response": response,
	})
}

func SignUpUserController(c echo.Context) error {
	var payloads = dto.UserRequest{}
	errBind := c.Bind(&payloads)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	emailExist := repository.CheckUserEmail(payloads.Email)
	if emailExist {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "fail sign up",
			"response": "email already exist",
		})
	}

	signUpData := dto.ConvertToUserModel(payloads)

	data, err := repository.CreateUser(signUpData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "fail sign up",
			"response": err.Error(),
		})
	}

	response := dto.ConvertToUserResponse(data)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success receive user data",
		"response": response,
	})
}

func ChangeUserPasswordController(c echo.Context) error {
	updatePassword := dto.ChangeUserPasswordRequest{}
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
	userData := dto.ConvertToChangeUserPasswordModel(updatePassword)

	responseData, err := repository.UpdateUserPassword(userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed change User",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success change password",
		"response": "success change password " + responseData.Email,
	})
}

func SendOTP(c echo.Context) error {
	var OTPReq = dto.OTPRequest{}
	errBind := c.Bind(&OTPReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	otp := fmt.Sprint(rand.Intn(99999 - 10000)+10000)

	err := repository.SetOTP(OTPReq.Email, otp)
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

func ValidateOTP(c echo.Context) error {
	var ValidateOTPReq = dto.ValidateOTPRequest{}
	errBind := c.Bind(&ValidateOTPReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	data, token, err := repository.ValidateOTP(ValidateOTPReq.Email, ValidateOTPReq.OTP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed validate otp",
			"response": err.Error(),
		})
	}
	response := map[string]any{
		"token":   token,
		"email":   data.Email,
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"message":  "Success validate otp",
		"response": response,
	})
}