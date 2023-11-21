package controller

import (
	"capstone-project/dto"
	"capstone-project/repository"
	"net/http"

	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "fail login",
			"response": err.Error(),
		})
	}
	response := map[string]any{
		"user_id": data.ID,
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

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success receive user data",
		"response": response,
	})
}

func ChangeUserPasswordController(c echo.Context) error {
	updatePassword := dto.UserRequest{}
	errBind := c.Bind(&updatePassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	emailExist := repository.CheckUserEmail(updatePassword.Email)
	if !emailExist {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "fail sign up",
			"response": "email doesn't exist",
		})
	}

	userData := dto.ConvertToUserModel(updatePassword)

	responseData, err := repository.UpdateUserPassword(userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed change User",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success change password",
		"response": "success change password for " + responseData.Email,
	})
}
