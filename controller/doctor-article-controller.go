package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllArticleDoctorsController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}
	responseData, err := repository.DoctorGetAllArticles(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get article data",
			"response": err.Error(),
		})
	}

	var articleResponse []dto.DoctorArticleResponse
	for _, article := range responseData {
		articleResponse = append(articleResponse, dto.ConvertToDoctorArticleResponse(article))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get article data",
		"response": articleResponse,
	})
}

func GetDoctorArticleByIDController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	articleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid article ID",
			"response": "Article ID must be a valid UUID.",
		})
	}

	article, err := repository.GetArticleByID(articleID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get article",
			"response": err.Error(),
		})
	}

	if article.DoctorID != doctor {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user's article.",
		})
	}

	articleResponse := dto.ConvertToDoctorArticleResponse(article)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get article",
		"response": articleResponse,
	})
}

func CreateDoctorArticleController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	article := dto.DoctorArticleRequest{}
	errBind := c.Bind(&article)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid body",
			"response": errBind.Error(),
		})
	}

	if len(article.Title) < 5 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid body",
			"response": "Title must be at least 5 characters long.",
		})
	}

	if article.Title == "" || len(article.Tags) == 0 || article.Reference == "" ||
		article.Image == "" || article.ImageDesc == "" || article.Content == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid body",
			"response": "All fields must be filled in.",
		})
	}

	articleData := dto.ConvertToDoctorArticleModel(article)
	articleData.DoctorID = doctor
	articleData.Published = false

	responseData, err := repository.InsertArticle(articleData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create article",
			"response": err.Error(),
		})
	}

	articleResponse := dto.ConvertToDoctorArticleResponse(responseData)

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create new article",
		"response": articleResponse,
	})
}

func UpdateDoctorArticleController(c echo.Context) error {
	articleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid article ID",
			"response": "Article ID must be a valid UUID.",
		})
	}

	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	article := dto.DoctorArticleRequest{}
	errBind := c.Bind(&article)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid body",
			"response": errBind.Error(),
		})
	}

	if len(article.Title) < 5 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid body",
			"response": "Title must be at least 5 characters long.",
		})
	}

	articleData, err := repository.GetArticleByID(articleID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed to get article",
			"response": err.Error(),
		})
	}

	if articleData.DoctorID != doctor {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to update other user's article.",
		})
	}

	articleData.Title = article.Title
	articleData.Tags = article.Tags
	articleData.Reference = article.Reference
	articleData.Image = article.Image
	articleData.ImageDesc = article.ImageDesc
	articleData.Content = article.Content

	err = repository.UpdateArticle(articleData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed to update article",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update article",
		"response": dto.ConvertToDoctorArticleResponse(articleData),
	})
}

func UpdateArticlePublishedStatusController(c echo.Context) error {
	articleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid article ID",
			"response": "Article ID must be a valid UUID.",
		})
	}

	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	err = repository.UpdateArticlePublishedStatus(articleID, doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed to update article status",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update article status",
		"response": articleID,
	})
}

func DeleteDoctorArticleController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	articleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid body",
			"response": err.Error(),
		})
	}

	article, err := repository.GetArticleByID(articleID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete article",
			"response": err.Error(),
		})
	}

	if article.DoctorID != doctor {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to delete other user's article.",
		})
	}

	err = repository.DeleteArticleByID(articleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed delete article",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete article",
		"response": "success delete article with id " + articleID.String(),
	})
}
