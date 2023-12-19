package controller_test

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/dto"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestGetBookmarkedArticlesController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get bookmarked articles",
			path:       "/articles/bookmarks",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetBookmarkedArticlesControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /articles/bookmarks", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetArticlesController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get articles",
			path:       "/articles",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetArticlesControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /articles/bookmarks", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestCreateCommentController(t *testing.T) {
	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	var testCases = []struct {
		name       string
		path       string
		patient    dto.CommentRequest
		expectCode int
	}{
		// {
		// 	name: "success create new comments",
		// 	path: "/articles/:id/comments",
		// 	patient: dto.CommentRequest{
		// 		PatientID: patient.ID,
		// 		Comment:   "Artikel yang sangat bermanfaat!",
		// 	},
		// 	expectCode: http.StatusCreated,
		// },
		{
			name: "failed create new comments",
			path: "/articles/:id/commentss",
			patient: dto.CommentRequest{
				PatientID: patient.ID,
				Comment:   "Artikel yang sangat bermanfaat!",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPost, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateCommentControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("POST /articles/:id/comments", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetArticleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get details article",
		// 	path:       "/articles/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get details article",
			path:       "/articless/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetArticleControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /articles/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestBookmarkController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get bookmark by id",
		// 	path:       "/articles/:id/bookmarks",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get bookmark by id",
			path:       "/articles/:id/bookmarks",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.BookmarkControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /articles/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
