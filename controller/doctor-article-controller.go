package controller

import (
	"capstone-project/dto"
	"capstone-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllArticleDotorsController(c echo.Context) error {
	responseData, err := repository.GetAllDoctorsArticles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get doctor articles",
			"response": err.Error(),
		})
	}

	var doctorArticleResponse []dto.DoctorArticleResponse
	for _, article := range responseData {
		doctorArticleResponse = append(doctorArticleResponse, dto.ConvertToDoctorArticleResponse(article))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get articles",
		"response": doctorArticleResponse,
	})
}

func CreateDoctorArticleController(c echo.Context) error {
	article := dto.DoctorArticleRequest{}
	errBind := c.Bind(&article)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	articleData := dto.ConvertToDoctorArticleModel(article)

	responseData, err := repository.InsertDoctorArticle(articleData)
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

func UpdateDoctorArticleController(c echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	updateArticleData := dto.DoctorArticleRequest{} // Ganti dengan tipe DTO yang sesuai untuk entitas dokter
	errBind := c.Bind(&updateArticleData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	doctorData := dto.ConvertToDoctorArticleModel(updateArticleData)

	responseData, err := repository.UpdateDoctorArticleByID(id, doctorData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed update doctor",
			"response": err.Error(),
		})
	}

	responseData, err = repository.GetDoctorArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get doctor",
			"reponse": err.Error(),
		})
	}

	doctorResponse := dto.ConvertToDoctorArticleResponse(responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update doctor",
		"response": doctorResponse,
	})
}

func DeleteDoctorArticleController(c echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	_, err = repository.GetDoctorArticleByID(id) // Ganti dengan repository yang sesuai
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete doctor",
			"reponse": err.Error(),
		})
	}

	err = repository.DeleteDoctorArticleByID(id) // Ganti dengan repository yang sesuai
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed delete doctor",
			"reponse": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success delete doctor",
		"response": "success delete doctor with id " + id.String(),
	})
}
