package controller

import (
	"capstone-project/dto"
	"capstone-project/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminLoginController(c echo.Context) error {
	var loginReq = dto.UserRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	data, token, err := repository.CheckAdmin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "fail login",
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