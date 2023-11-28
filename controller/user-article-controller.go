package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"
	m "capstone-project/middleware"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetArticlesController(c echo.Context) error {
	responseData, err := repository.GetAllArticleDashboard()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get articles",
			"response":   err.Error(),
		})
	}

	var articleResponse []dto.DoctorArticleResponse
	for _, article := range responseData {
		articleResponse = append(articleResponse, dto.ConvertToDoctorArticleResponse(article))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get articles",
		"response":   articleResponse,
	})
}

func CreateCommentController(c echo.Context) error {
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

	_, err = repository.GetArticleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "failed create comment",
			"response": err.Error(),
		})
	}

	comment := dto.CommentRequest{}
	errBind := c.Bind(&comment)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(comment.PatientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create comment",
			"reponse":   err.Error(),
		})
	}
	if checkPatient.UserID != user{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user patient data.",
		})
	}

	commentData := dto.ConvertToCommentModel(comment)
	commentData.ArticleID = uuid
	
	responseData, err := repository.InsertComment(commentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create comment",
			"response":  err.Error(),
		})
	}

	commentResponse := dto.ConvertToCommentResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new comment",
		"response":    commentResponse,
	})
}

func GetArticleController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.UserGetArticleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get article",
			"reponse": err.Error(),
		})
	}

	articleResponse := dto.ConvertToUserArticleResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get article",
		"response": articleResponse,
	})
}