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
	doctor_id := c.FormValue("doctor_id")

	responseData, err := repository.GetAllArticles(doctor_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get article",
			"response": err.Error(),
		})
	}

	var articleResponse []dto.DoctorArticleResponse
	for _, article := range responseData {
		articleResponse = append(articleResponse, dto.ConvertToDoctorArticleResponse(article))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get article",
		"response": articleResponse,
	})
}

func CreateDoctorArticleController(c echo.Context) error {
	doctor, err := m.ExtractTokenDoctorId(c)
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

	checkDoctor, err := repository.GetDoctorByID(article.DoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create article",
			"response": err.Error(),
		})
	}
	if checkDoctor.ID != doctor {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: You are not allowed to access other user doctor data.",
		})
	}

	articleData := dto.ConvertToDoctorArticleModel(article)

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
	doctor, err := m.ExtractTokenDoctorId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: Doctor is not valid.",
		})
	}

	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	checkArticle, err := repository.GetArticleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete article",
			"reponse": err.Error(),
		})
	}

	checkDoctor, err := repository.GetArticleByID(checkArticle.DoctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete article",
			"reponse": err.Error(),
		})
	}
	if checkDoctor.DoctorID != doctor {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "unauthorized",
			"reponse": "Permission Denied: You are not allowed to access other user doctor data.",
		})
	}

	err = repository.DeleteArticleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed delete doctor",
			"reponse": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor",
		"response": "success delete doctor with id " + uuid.String(),
	})
}
