package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	"capstone-project/model"
	"capstone-project/middleware"
	"capstone-project/constant"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func InsertDataUser() (string) {
	user := model.User{
		ID: uuid.New(),
		Name:         "Davin2",
		Email:        "davin2@gmail.com",
		Password:     "12345678",
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	database.DB.Create(&user)

	token, _ := middleware.CreateToken(user.ID, constant.ROLE_USER, user.Name, false)
	return token
}

func TestSignUpController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user	   model.User
		expectCode int
	}{
		{
			name:       "create new user",
			path:       "/users",
			user:		model.User{
							Name:     "Davin2",
							Password: "12345678",
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusCreated,
		},
		{
			name:       "create new user email existed",
			path:       "/users",
			user:		model.User{
							Name:     "Davin2",
							Password: "12345678",
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "create new user invalid password",
			path:       "/users",
			user:		model.User{
							Name:     "Davin Error",
							Password: "123",
							Email:    "davinError@gmail.com",
						},
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "create new user invalid email",
			path:       "/users",
			user:		model.User{
							Name:     "Davin Error",
							Password: "12345678",
							Email:    "davinError",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)
		
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.SignUpUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Response   dto.UserResponse    		`json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Response.Name, testCase.user.Name)
				assert.Equal(t, responseData.Response.Email, testCase.user.Email)
			}
		}
	}
}

func TestLoginController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user	   model.User
		expectCode int
	}{
		{
			name:       "Login",
			path:       "/users/login",
			user:		model.User{
							Password: "12345678",
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "Login",
			path:       "/users/login",
			user:		model.User{
							Password: "12345678",
							Email:    "Invalid email",
						},
			expectCode: http.StatusUnauthorized,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUser()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)
		
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.LoginUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type LoginResponse struct {
				UserID  string    `json:"user_id"`
				Email string    `json:"email"`
				Name string    `json:"name"`
				Token string    `json:"token"`
			}
			type Response struct {
				Message string           `json:"message"`
				Response   LoginResponse `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Response.Email, testCase.user.Email)
			}
		}
	}
}

func TestSendOTPController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user	   model.User
		expectCode int
	}{
		{
			name:       "Send OTP",
			path:       "/users/send-otp",
			user:		model.User{
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "Send OTP",
			path:       "/users/send-otp",
			user:		model.User{
							Email:    "davin@gmail.com",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUser()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)
		
		req := httptest.NewRequest(http.MethodPut, "/users/send-otp", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.SendOTP(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message string           `json:"message"`
				Response   map[string]any `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestValidateOTPController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user	   model.User
		expectCode int
	}{
		{
			name:       "Validate OTP",
			path:       "/users/validate-otp",
			user:		model.User{
							Email:    "davin2@gmail.com",
							OTP:    "",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "Validate OTP Wrong OTP",
			path:       "/users/validate-otp",
			user:		model.User{
							Email:    "davin2@gmail.com",
							OTP:    "1",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUser()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)
		
		req := httptest.NewRequest(http.MethodPut, "/users/validate-otp", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.ValidateOTP(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message string           `json:"message"`
				Response   map[string]any `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}