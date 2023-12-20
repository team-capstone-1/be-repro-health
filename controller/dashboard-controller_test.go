package controller_test

import (
	"capstone-project/config"
	"capstone-project/controller"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestGetDataCountForDoctorControllerOneMonth(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "Failed to get data count for doctor one month",
			path:       "/dashboard/data-count-one-month",
			expectCode: http.StatusBadRequest,
		},
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

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDataCountForDoctorControllerOneMonthTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetDataCountForDoctorControllerOneMonth_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/dashboard/data-count-one-month"+tokenString, nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/dashboard/data-count-one-month" + tokenString)

	c.Set("user", token)

	err = controller.GetDataCountForDoctorControllerOneMonth(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestGetDataCountForDoctorControllerOneWeek(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "Failed to get data count for doctor one week",
			path:       "/dashboard/data-count-one-week",
			expectCode: http.StatusBadRequest,
		},
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

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDataCountForDoctorControllerOneWeekTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetDataCountForDoctorControllerOneWeek_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/dashboard/data-count-one-month"+tokenString, nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/dashboard/data-count-one-month" + tokenString)

	c.Set("user", token)

	err = controller.GetDataCountForDoctorControllerOneWeek(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestGetDataCountForDoctorControllerOneDay(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "Failed to get data count for doctor one day",
			path:       "/dashboard/data-count-one-day",
			expectCode: http.StatusBadRequest,
		},
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

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDataCountForDoctorControllerOneDayTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetDataCountForDoctorControllerOneDay_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/dashboard/data-count-one-month"+tokenString, nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/dashboard/data-count-one-month" + tokenString)

	c.Set("user", token)

	err = controller.GetDataCountForDoctorControllerOneDay(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
