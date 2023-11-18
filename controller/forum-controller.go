package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

func GetForumsController(c echo.Context) error {
	title := c.FormValue("title")

	responseData, err := repository.GetAllForums(title)
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

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get forums",
		"response":   forumResponse,
	})
}

func CreateForumController(c echo.Context) error {
	forum := dto.ForumRequest{}
	errBind := c.Bind(&forum)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
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

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new forum",
		"response":    forumResponse,
	})
}

func DeleteForumController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	_, err = repository.GetForumByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete forum",
			"reponse":   err.Error(),
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
