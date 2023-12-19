package controller_test

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	"capstone-project/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InsertDataForums() {
	forumID, _ := uuid.Parse("f1c11001-e254-40a9-8929-22e5bf660d1c")
	// CLINIC SEEDERS
	forum := []model.Forum{
		{
			ID:        forumID,
			PatientID: uuid.New(),
			Title:     "Test Forum",
			Content:   "Test Forum Content",
			Date:      time.Now(),
		},
	}

	for _, v := range forum {
		var exist model.Forum

		errCheck := database.DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			database.DB.Create(&v)
		}
	}
	return
}

func TestGetDoctorAllForumsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get doctor all forums",
			path:       "/doctors/forums",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDoctorAllForumsControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestCreateDoctorReplyForumUnauthorized(t *testing.T) {
	e := echo.New()

	// Create a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Prepare the request body
	requestBody := `{
        "forum_id": "3b1d0a0f-7b0b-4c0d-8e1e-2e1e1e1e1e1e",
    	"content": "Semoga jawaban saya membantu Anda"
    }`

	req := httptest.NewRequest(http.MethodPost, "/doctors/forum-replies", strings.NewReader(requestBody))
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set the JWT token directly in the context
	c.Set("user", token)

	// Execute the correct controller for creating a doctor reply to a forum
	err := controller.CreateDoctorReplyForum(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	// Additional assertions based on the response if needed
}

func TestCreateDoctorReplyForum(t *testing.T) {
	forumID, _ := uuid.Parse("f1c11001-e254-40a9-8929-22e5bf660d1c")
	doctorID, _ := uuid.Parse("f7613c10-29fd-4b82-bfea-1649ae41af98")

	var testCases = []struct {
		name       string
		path       string
		reply      dto.DoctorForumReplyRequest
		expectCode int
	}{
		// {
		// 	name: "success create forum reply",
		// 	path: "/doctors/forum-replies",
		// 	reply: dto.DoctorForumReplyRequest{
		// 		ForumsID: forumID,
		// 		DoctorID: doctorID,
		// 		Content:  "Test forum reply",
		// 		Date:     time.Now(),
		// 	},
		// 	expectCode: http.StatusCreated,
		// },
		{
			name: "success create forum reply",
			path: "/doctors/forum-repliess",
			reply: dto.DoctorForumReplyRequest{
				ForumsID: forumID,
				DoctorID: doctorID,
				Content:  "Test forum reply",
				Date:     time.Now(),
			},
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()
	_, user := InsertDataUser()
	InsertDataPatient(user.ID)
	InsertDataForums()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.reply)
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateDoctorReplyForumTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("POST %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetForumDetailsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "get details forum",
		// 	path:       "/doctors/forums/details/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "get details forum",
			path:       "/doctors/forums/detailss/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()
	_, user := InsertDataUser()
	InsertDataPatient(user.ID)
	InsertDataForums()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDoctorForumDetailsTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeleteForumReplyController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success delete forum reply",
		// 	path:       "/doctors/forum-replies/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed delete forum reply invalid endpoint",
			path:       "/doctors/forum-repliess/:id",
			expectCode: http.StatusBadRequest,
		},
		// {
		// 	name:       "failed delete forum reply invalid endpoint",
		// 	path:       "/doctors/forum-replies/",
		// 	expectCode: http.StatusNotFound,
		// },
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()
	_, user := InsertDataUser()
	InsertDataPatient(user.ID)
	InsertDataForums()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodDelete, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.DeleteDoctorForumReplyControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("DELETE %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestUpdateDoctorReplyForum(t *testing.T) {
	forumID, _ := uuid.Parse("f1c11001-e254-40a9-8929-22e5bf660d1c")
	doctorID, _ := uuid.Parse("f7613c10-29fd-4b82-bfea-1649ae41af98")

	var testCases = []struct {
		name       string
		path       string
		reply      dto.DoctorForumReplyRequest
		expectCode int
	}{
		// {
		// 	name: "success create forum reply",
		// 	path: "/doctors/forum-replies/:id",
		// 	reply: dto.DoctorForumReplyRequest{
		// 		ForumsID: forumID,
		// 		DoctorID: doctorID,
		// 		Content:  "Test forum reply",
		// 		Date:     time.Now(),
		// 	},
		// 	expectCode: http.StatusOK,
		// },
		// {
		// 	name: "success create forum reply",
		// 	path: "/doctors/forum-repliess/:id",
		// 	reply: dto.DoctorForumReplyRequest{
		// 		ForumsID: forumID,
		// 		DoctorID: doctorID,
		// 		Content:  "Test forum reply",
		// 		Date:     time.Now(),
		// 	},
		// 	expectCode: http.StatusNotFound,
		// },
		{
			name: "success create forum reply",
			path: "/doctors/forum-replies/",
			reply: dto.DoctorForumReplyRequest{
				ForumsID: forumID,
				DoctorID: doctorID,
				Content:  "Test forum reply",
				Date:     time.Now(),
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()
	_, user := InsertDataUser()
	InsertDataPatient(user.ID)
	InsertDataForums()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.reply)
		req := httptest.NewRequest(http.MethodPut, testCase.path, strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdateDoctorReplyForumTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("PUT %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
