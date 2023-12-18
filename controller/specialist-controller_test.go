package controller_test

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"fmt"
	"io"
	"os"

	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	"capstone-project/model"
	"capstone-project/config"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InsertDataSpecialistForGetSpecialists() (model.Specialist, error) {
	specialist := model.Specialist{
		ID: uuid.New(),
		Name:         "Sample Specialist",
		Image:        "sample_image.jpg",
	}

	var err error
	if err = database.DB.Create(&specialist).Error; err != nil {
		return model.Specialist{}, err
	}
	return specialist, nil
}

func InsertDataSpecialistForGetSpecialistsByClinicID() uuid.UUID {
	specialistID := uuid.New()
	clinicID := uuid.New()
	doctorID := uuid.New()

	specialist := model.Specialist{
		ID: specialistID,
	}

	clinic := model.Clinic{
		ID: clinicID,
	}

	doctor := model.Doctor{
		ID: doctorID,
		ClinicID: clinicID,
		SpecialistID: specialistID,
	}

	database.DB.Create(&specialist)
	database.DB.Create(&clinic)
	database.DB.Create(&doctor)

	return clinicID
}

func TestGetSpecialistsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all specialist",
			path:       "/specialists",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataSpecialistForGetSpecialists()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, controller.GetSpecialistsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message string                   `json:"message"`
				Responses   []dto.SpecialistResponse 	 `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(responseData.Responses))
		}
	}
}

func TestCreateSpecialistController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		specialist dto.SpecialistRequest
		expectCode int
	}{
		{
			name:       "create new specialist",
			path:       "/admins/specialists",
			specialist:		dto.SpecialistRequest{
								Name: "Kandungan",
							},
			expectCode: http.StatusCreated,
		},
	}

	e := InitEchoTestAPI()
	token := InsertDataAdmin()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "your_image_file.jpg")
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
		
		req := httptest.NewRequest(http.MethodPost, "/admins/specialists", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateSpecialistControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("POST /admins/specialists", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestUpdateSpecialistController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		specialist dto.SpecialistRequest
		expectCode int
	}{
		{
			name:       "create new specialist",
			path:       "/admins/specialists/:id",
			specialist:		dto.SpecialistRequest{
								Name: "Kandungan",
							},
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token := InsertDataAdmin()
	specialist, _ := InsertDataSpecialistForGetSpecialists()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "your_image_file.jpg")
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
		
		req := httptest.NewRequest(http.MethodPut, "/admins/specialists/:id", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues((specialist.ID).String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdateSpecialistControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("PUT /admins/specialists/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeleteSpecialistByIDController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get specialists",
			path:       "/specialists/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	token := InsertDataAdmin()
	specialist, _ := InsertDataSpecialistForGetSpecialists()

	for _, testCase := range testCases {
		
		req := httptest.NewRequest(http.MethodDelete, "/admins/specialists/:id", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(specialist.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DeleteSpecialistControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("DELETE /admins/specialists/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetSpecialistsByClinicIDController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all specialist",
			path:       "/clinics/:id/specialists",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	clinicID := InsertDataSpecialistForGetSpecialistsByClinicID()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(clinicID.String())

		if assert.NoError(t, controller.GetSpecialistsByClinicController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message string                   `json:"message"`
				Responses   []dto.SpecialistResponse 	 `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(responseData.Responses))
		}
	}
}