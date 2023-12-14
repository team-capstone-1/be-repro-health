package controller

import (
	"capstone-project/dto"
	"capstone-project/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminLoginController(c echo.Context) error {
	var loginReq = dto.LoginRequest{}
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

	data, token, err := repository.CheckAdmin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "admin account not found",
			"response": err.Error(),
		})
	}

	response := map[string]any{
		"admin_id": data.ID,
		"email":    data.Email,
		"token":    token,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success login admin account",
		"response": response,
	})
}
