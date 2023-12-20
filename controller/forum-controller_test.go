package controller_test

import (
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/dto"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

// func TestGetForumsController(t *testing.T) {
// 	// Initialize Echo
// 	e := echo.New()

// 	// Inisialisasi objek controller
// 	// controller := &ForumController{}

// 	req := httptest.NewRequest(http.MethodGet, "/forums", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	c.SetPath("/forums")

// 	// Call the controller function
// 	err := controller.GetForumsController(c)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// }

func TestGetForumController(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Create a new request with a valid forum ID
	forumID := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/forums/"+forumID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the controller function
	err := controller.GetForumController(c)

	c.SetPath("/forums" + forumID.String())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateForumController(t *testing.T) {

	e := InitEchoTestAPI()
	token, _ := InsertDataUser()
	// patient, _ := InsertDataPatient(user.ID)

	var testCases = []struct {
		name       string
		path       string
		patient    dto.ForumRequest
		expectCode int
	}{
		// {
		// 	name: "success create forum",
		// 	path: "/forums",
		// 	patient: dto.ForumRequest{
		// 		PatientID: patient.ID,
		// 		Title:     "Saya sakit perut",
		// 		Content:   "Saat malam saya sering mengalami sakit perut bagian bawah",
		// 		Anonymous: true,
		// 	},
		// 	expectCode: http.StatusOK,
		// },
		{
			name: "failed create forum",
			path: "/forums",
			patient: dto.ForumRequest{
				PatientID: uuid.Nil,
				Title:     "Saya sakit perut",
				Content:   "Saat malam saya sering mengalami sakit perut bagian bawah",
				Anonymous: true,
			},
			expectCode: http.StatusBadRequest,
		},
		// {
		// 	name: "failed create forum invalid endpoint",
		// 	path: "/forumss",
		// 	patient: dto.ForumRequest{
		// 		PatientID: patient.ID,
		// 		Title:     "Saya sakit perut",
		// 		Content:   "Saat malam saya sering mengalami sakit perut bagian bawah",
		// 		Anonymous: true,
		// 	},
		// 	expectCode: http.StatusNotFound,
		// },
	}

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPost, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		// context.SetParamNames("id")
		// context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateForumControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("POST %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeleteForumController(t *testing.T) {

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	var testCases = []struct {
		name       string
		path       string
		patient    dto.ForumRequest
		expectCode int
	}{
		// {
		// 	name: "success delete forum",
		// 	path: "/forums/:id",
		// 	patient: dto.ForumRequest{
		// 		PatientID: patient.ID,
		// 		Title:     "Saya sakit perut",
		// 		Content:   "Saat malam saya sering mengalami sakit perut bagian bawah",
		// 		Anonymous: true,
		// 	},
		// 	expectCode: http.StatusOK,
		// },
		{
			name: "failed delete forum",
			path: "/forums/:id",
			patient: dto.ForumRequest{
				PatientID: uuid.Nil,
				Title:     "Saya sakit perut",
				Content:   "Saat malam saya sering mengalami sakit perut bagian bawah",
				Anonymous: true,
			},
			expectCode: http.StatusBadRequest,
		},
		// {
		// 	name: "failed delete forum invalid endpoint",
		// 	path: "/forumss/:id",
		// 	patient: dto.ForumRequest{
		// 		PatientID: patient.ID,
		// 		Title:     "Saya sakit perut",
		// 		Content:   "Saat malam saya sering mengalami sakit perut bagian bawah",
		// 		Anonymous: true,
		// 	},
		// 	expectCode: http.StatusNotFound,
		// },
	}

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodDelete, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DeleteForumControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("DELETE %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
