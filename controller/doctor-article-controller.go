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
	responseData, err := repository.GetAllArticles()
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
		"message":  "success get transaction data",
		"response": articleResponse,
	})
}

func CreateDoctorArticleController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	article := dto.DoctorArticleRequest{}
	errBind := c.Bind(&article)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	articleData := dto.ConvertToDoctorArticleModel(article)
	articleData.DoctorID = doctor

	responseData, err := repository.InsertArticle(articleData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create article",
			"response": err.Error(),
		})
	}

	articleResponse := dto.ConvertToDoctorArticleResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create new article",
		"response": articleResponse,
	})
}

func DeleteDoctorArticleController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	checkArticle, err := repository.GetArticleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete article",
			"response": err.Error(),
		})
	}

	checkDoctor, err := repository.GetDoctorByID(checkArticle.DoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete article",
			"response": err.Error(),
		})
	}
	if checkDoctor.ID != doctor {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor data.",
		})
	}

	err = repository.DeleteArticleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed delete doctor",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success delete doctor",
		"response": "success delete doctor with id " + uuid.String(),
	})
}