package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"time"
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
		UserID: user,
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
		patient	   model.Patient
		expectCode int
	}{
		{
			name:       "create new patient",
			path:       "/patients",
			patient:		model.Patient{
							Name:            "Davin2",
							TelephoneNumber: "123456789",
							DateOfBirth:     time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
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
		patientJSON, _ := json.Marshal(testCase.patient)
		
		req := httptest.NewRequest(http.MethodPost, "/patients", strings.NewReader(string(patientJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set(echo.HeaderAuthorization, "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJuYW1lIjoiRGF2aW5ubiIsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjUwMWM3MzdhLTcyY2EtNGY1ZS04YjM1LWY1Mzc0ZTRmZDg1YyJ9.Ioa0l1n0vJpqi0BrQOWT0skSEMMGxi49g_y3_QrBh0w")
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreatePatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.CreatePatientController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Response   dto.PatientResponse    		`json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)
			fmt.Println("token:", responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Response.Name, testCase.patient.Name)
			}
		}
	}
}