package controller_test

import (
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	"capstone-project/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InsertDataPatient(userid uuid.UUID) (model.Patient, error) {
	// user := m.ExtractTokenUserId(c)
	patientID, _ := uuid.Parse("d4f1bd00-fbd5-4b04-93f8-aa889d15ad5a")
	patient := model.Patient{
		ID:              patientID,
		Name:            "Davin2",
		UserID:          userid,
		ProfileImage:    "",
		TelephoneNumber: "123456789",
		DateOfBirth:     time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Relation:        "sibling",
		Weight:          70.5,
		Height:          175.0,
		Gender:          "male",
	}

	var err error
	if err = database.DB.Create(&patient).Error; err != nil {
		return model.Patient{}, err
	}
	return patient, nil
}

func TestCreatePatientController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		patient    dto.PatientRequest
		expectCode int
	}{
		{
			name: "create new patient",
			path: "/patients",
			patient: dto.PatientRequest{
				Name:            "Davin2",
				TelephoneNumber: "123456789",
				DateOfBirth:     time.Now(),
				Relation:        "sibling",
				Weight:          70.5,
				Height:          175.0,
				Gender:          "male",
			},
			expectCode: http.StatusCreated,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataUser()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("profile_image", "your_image_file.jpg")
		if err != nil {
			t.Fatal(err)
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/patients", &buf)
		// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreatePatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("POST /patients", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetAllPatientsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get patient",
			path:       "/patients",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, "/patients/", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetPatientsControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /patients", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestUpdatePatientController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		patient    dto.PatientRequest
		expectCode int
	}{
		{
			name: "update patient",
			path: "/patients/:id",
			patient: dto.PatientRequest{
				Name:            "Davin2",
				TelephoneNumber: "123456789",
				DateOfBirth:     time.Now(),
				Relation:        "sibling",
				Weight:          70.5,
				Height:          175.0,
				Gender:          "male",
			},
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("profile_image", "your_image_file.jpg")
		if err != nil {
			t.Fatal(err)
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/patients/:id", &buf)
		// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdatePatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("PUT /patients/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetPatientByIDController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get patient",
			path:       "/patients/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, "/patients/:id", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetPatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /patients/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetPatientByIDControllerInvalid(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get patient",
			path:       "/patients/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, "/patients/:id", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.GetPatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /patients/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeletePatientByIDController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get patient",
			path:       "/patients/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodDelete, "/patients/:id", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DeletePatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("DELETE /patients/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
