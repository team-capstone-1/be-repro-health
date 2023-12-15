package controller

import (
	"net/http"

	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetBookmarkedArticlesController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	userBookmark, err := repository.GetArticleBookmark(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get bookmarks",
			"response": err.Error(),
		})
	}

	var articleResponse []dto.UserArticleResponse
	for _, article := range userBookmark {
		articleResponse = append(articleResponse, dto.ConvertToUserArticleResponse(article, true))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get bookmarked articles",
		"response": articleResponse,
	})
}

func GetArticlesController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	responseData, err := repository.UserGetAllArticle()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get articles",
			"response": err.Error(),
		})
	}

	userBookmark, err := repository.GetArticleBookmark(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get bookmarks",
			"response": err.Error(),
		})
	}

	bookmarkMap := make(map[uuid.UUID]struct{})
	for _, bookmark := range userBookmark {
		bookmarkMap[bookmark.ID] = struct{}{}
	}

	var articleResponse []dto.UserArticleResponse
	for _, article := range responseData {
		_, bookmarked := bookmarkMap[article.ID]
		articleResponse = append(articleResponse, dto.ConvertToUserArticleResponse(article, bookmarked))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get articles",
		"response": articleResponse,
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
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID(comment.PatientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create comment",
			"reponse": err.Error(),
		})
	}
	if checkPatient.UserID != user {
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
			"message":  "failed create comment",
			"response": err.Error(),
		})
	}

	commentResponse := dto.ConvertToCommentResponse(responseData, repository.GetProfileByPatientID(responseData.PatientID))

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success create new comment",
		"response": commentResponse,
	})
}

func GetArticleController(c echo.Context) error {
	user := m.ExtractTokenUserId(c)
	if user == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Permission Denied: User is not valid.",
		})
	}

	articleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	responseData, err := repository.UserGetArticleByID(articleID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get article",
			"reponse": err.Error(),
		})
	}

	userBookmark, err := repository.GetArticleBookmark(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get bookmarks",
			"response": err.Error(),
		})
	}

	bookmarkMap := make(map[uuid.UUID]struct{})
	for _, bookmark := range userBookmark {
		bookmarkMap[bookmark.ID] = struct{}{}
	}

	_, bookmarked := bookmarkMap[articleID]
	articleResponse := dto.ConvertToUserArticleResponse(responseData, bookmarked)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get article",
		"response": articleResponse,
	})
}


func BookmarkController(c echo.Context) error {
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

	err = repository.UpdateBookmark(user, uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update bookmark",
			"response": err.Error(),
		})
	}

	userBookmark, err := repository.GetArticleBookmark(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get bookmarks",
			"response": err.Error(),
		})
	}

	var articleResponse []dto.UserArticleResponse
	for _, article := range userBookmark {
		articleResponse = append(articleResponse, dto.ConvertToUserArticleResponse(article, true))
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message":  "success update bookmark",
		"response": articleResponse,
	})
}