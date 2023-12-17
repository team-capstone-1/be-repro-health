package controller_test

import (
	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"mime/multipart"
	"testing"
	// "strings"
	"time"
	"io"
	"os"
	"bytes"
	"fmt"

	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	"capstone-project/model"
	"capstone-project/config"
	m "capstone-project/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4/middleware"
)

func InsertDataPatient(c echo.Context) (model.Patient, error) {
	user := m.ExtractTokenUserId(c)
	patient := model.Patient{
		Name:            "Davin2",
		UserID: 		 user,
		ProfileImage:	 "",
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
		patient	   dto.PatientRequest
		expectCode int
	}{
		{
			name:       "create new patient",
			path:       "/patients",
			patient:		dto.PatientRequest{
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
	token := InsertDataUser()

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