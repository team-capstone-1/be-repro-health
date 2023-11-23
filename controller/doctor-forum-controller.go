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

	forum := dto.DoctorForumReplyRequest{DoctorID: user}
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

func UpdateDoctorReplyForum(c echo.Context) error {
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
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	updateData := dto.DoctorUpdateForumReplyRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	forumReply := dto.ConvertToDoctorUpdateForumReplyModel(updateData)

	responseData, err := repository.UpdateDoctorReplyForum(uuid, forumReply)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed update forum reply",
			"response": err.Error(),
		})
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	responseData, err = repository.GetDoctorForumReplyByID(uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "failed get forum reply",
			"reponse": err.Error(),
		})
	}

	forumReplyResponse := dto.ConvertToDoctorForumReplyResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update forum reply",
		"response": forumReplyResponse,
	})
}

func GetDoctorForumReplyID(c echo.Context) error {
	// BUTUH GET FORUM REPLY ID SUPAYA BISA UPDATE

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
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	//recall the GetById repo because if I return it from update, it only fill the updated field and leaves everything else null or 0
	responseData, err := repository.GetDoctorForumReplyByID(uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "failed get forum reply",
			"reponse": err.Error(),
		})
	}

	patientResponse := dto.ConvertToDoctorForumReplyResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update forum reply",
		"response": patientResponse,
	})
}

func DeleteDoctorForumReplyController(c echo.Context) error {
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
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	checkForum, err := repository.GetDoctorForumReplyByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete forum reply",
			"reponse": err.Error(),
		})
	}

	checkDoctor, err := repository.GetDoctorByID(checkForum.DoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete forum",
			"reponse": err.Error(),
		})
	}
	if checkDoctor.ID != user {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	err = repository.DeleteForumReplyByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed delete forum reply",
			"reponse": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete forum reply",
		"response": "success delete forum reply with id " + uuid.String(),
	})
}
