package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetDoctorAllForumsController(c echo.Context) error {
	title := c.FormValue("title")
	patient_id := c.FormValue("patient_id")

	responseData, err := repository.DoctorGetAllForums(title, patient_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get forums",
			"response": err.Error(),
		})
	}

	var forumResponse []dto.DoctorForumResponse
	for _, forum := range responseData {
		forumResponse = append(forumResponse, dto.ConvertToDoctorForumResponse(forum))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get forums",
		"response": forumResponse,
	})
}

func CreateDoctorReplyForum(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	forum := dto.DoctorForumReplyRequest{}
	errBind := c.Bind(&forum)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	forumData := dto.ConvertToDoctorReplyModel(forum)

	responseData, err := repository.CreateDoctorReplyForum(forumData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create forum reply",
			"response": err.Error(),
		})
	}

	forumResponse := dto.ConvertToDoctorForumReplyResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create forum reply",
		"response": forumResponse,
	})
}
