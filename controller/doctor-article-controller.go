package controller

import (
	"capstone-project/dto"
	"capstone-project/model"
	"capstone-project/repository"
	"capstone-project/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DoctorArticleController struct {
	Repo *repository.DoctorArticleRepository
}

func NewDoctorArticleController(db *gorm.DB) *DoctorArticleController {
	return &DoctorArticleController{
		Repo: repository.NewDoctorArticleRepository(db),
	}
}

func (c *DoctorArticleController) GetListAllArticlesDoctor(ctx echo.Context) error {
	articles, err := c.Repo.GetAllArticles()
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusInternalServerError, "Internal Server Error")
		return err
	}

	responseList := make([]*dto.DoctorArticleResponse, len(articles))
	for i, article := range articles {
		responseList[i] = dto.ConvertToDoctorArticleResponse(&article)
	}

	return ctx.JSON(http.StatusOK, responseList)
}

func (c *DoctorArticleController) GetArticleByIDDoctor(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusBadRequest, "Invalid ID")
		return err
	}

	article, err := c.Repo.GetArticleByID(uint(id))
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusNotFound, "Article not found")
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}

func (c *DoctorArticleController) CreateNewArticleDoctor(ctx echo.Context) error {
	var request dto.DoctorArticleRequest
	if err := ctx.Bind(&request); err != nil {
		util.RespondWithError(ctx.Response(), http.StatusBadRequest, "Invalid request payload")
		return err
	}

	doctorID := uint(1) // Placeholder, replace with actual authentication logic

	articleModel := &model.Article{
		DoctorID: doctorID,
		Title:    request.Title,
		Content:  request.Content,
		Image:    request.Image,
	}

	err := c.Repo.CreateArticle(articleModel)
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusInternalServerError, "Internal Server Error")
		return err
	}

	return ctx.JSON(http.StatusCreated, articleModel)
}

func (c *DoctorArticleController) UpdateArticleByIdDoctor(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusBadRequest, "Invalid ID")
		return err
	}

	article, err := c.Repo.GetArticleByID(uint(id))
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusNotFound, "Article not found")
		return err
	}

	var request dto.DoctorArticleRequest
	if err := ctx.Bind(&request); err != nil {
		util.RespondWithError(ctx.Response(), http.StatusBadRequest, "Invalid request payload")
		return err
	}

	article.Title = request.Title
	article.Content = request.Content
	article.Image = request.Image

	err = c.Repo.UpdateArticle(article)
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusInternalServerError, "Internal Server Error")
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}

func (c *DoctorArticleController) DeleteArticleByIdDoctor(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusBadRequest, "Invalid ID")
		return err
	}

	err = c.Repo.DeleteArticle(uint(id))
	if err != nil {
		util.RespondWithError(ctx.Response(), http.StatusInternalServerError, "Internal Server Error")
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}