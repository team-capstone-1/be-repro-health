package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"
	m "capstone-project/middleware"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

func GetForumsController(c echo.Context) error {
	title := c.FormValue("title")
	patient_id := c.FormValue("patient_id")

	responseData, err := repository.GetAllForums(title, patient_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get forums",
			"response":   err.Error(),
		})
	}

	var forumResponse []dto.ForumResponse
	for _, forum := range responseData {
		forumResponse = append(forumResponse, dto.ConvertToForumResponse(forum))
	}

	for _, forumRes := range forumResponse{
		forumRes.Profile = repository.GetProfileByPatientID(forumRes.PatientID)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get forums",
		"response":   forumResponse,
	})
}

func GetForumController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.GetForumByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get forum",
			"reponse": err.Error(),
		})
	}

	forumResponse := dto.ConvertToForumResponse(responseData)
	forumResponse.Profile = repository.GetProfileByPatientID(forumResponse.PatientID)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get forum",
		"response": forumResponse,
	})
}

func CreateForumController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	forum := dto.ForumRequest{}
	errBind := c.Bind(&forum)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(forum.PatientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create forum",
			"reponse":   err.Error(),
		})
	}
	if checkPatient.UserID != user{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	forumData := dto.ConvertToForumModel(forum)
	
	responseData, err := repository.InsertForum(forumData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create forum",
			"response":  err.Error(),
		})
	}

	forumResponse := dto.ConvertToForumResponse(responseData)
	forumResponse.Profile = repository.GetProfileByPatientID(forumResponse.PatientID)

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "success create new forum",
		"response":    forumResponse,
	})
}

func DeleteForumController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	checkForum, err := repository.GetForumByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete forum",
			"reponse":   err.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(checkForum.PatientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete forum",
			"reponse":   err.Error(),
		})
	}
	if checkPatient.UserID != user{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	err = repository.DeleteForumByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed delete forum",
			"reponse":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete forum",
		"response": "success delete forum with id " + uuid.String(),
	})
}
