package controller_test

import (
	"capstone-project/config"
	"capstone-project/controller"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestGetAllDoctorScheduleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get all schedules",
			path:       "/doctors/schedule",
			expectCode: http.StatusOK,
		},
		// {
		// 	name:       "failed get all schedules",
		// 	path:       "/doctors/schedules",
		// 	expectCode: http.StatusBadRequest,
		// },
	}

	e := InitEchoTestAPI()
	token, doctor := InsertDataDoctor()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetAllDoctorScheduleControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /doctors/schedule", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDoctorInactiveScheduleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success inactive doctor",
			path:       "/doctors/schedule/inactive?date=2023-12-21&session=pagi",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, doctor := InsertDataDoctor()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPut, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorInactiveScheduleControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("PUT /doctors/schedule/inactive?date=2023-12-21&session=pagi", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDoctorActiveScheduleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success active doctor",
			path:       "/doctors/schedule/active?date=2023-12-21&session=pagi",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, doctor := InsertDataDoctor()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPut, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorActiveScheduleControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("PUT /doctors/schedule/active?date=2023-12-21&session=pagi", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
