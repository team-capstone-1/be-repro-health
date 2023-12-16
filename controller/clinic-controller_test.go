package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	"capstone-project/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	database.InitTest()
	e := echo.New()
	return e
}

func InsertDataClinicForGetClinics() (model.Clinic, error) {
	clinic := model.Clinic{
		Name:         "Sample Clinic",
		Image:        "sample_image.jpg",
		City:         "Sample City",
		Location:     "Sample Location",
		Telephone:    "123456789",
		Email:        "sample@example.com",
		Profile:      "Sample Profile",
		Latitude:     "12.345",
		Longitude:    "67.890",
	}

	var err error
	if err = database.DB.Create(&clinic).Error; err != nil {
		return model.Clinic{}, err
	}
	return clinic, nil
}

func TestGetClinicsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all clinic",
			path:       "/clinics",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataClinicForGetClinics()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, controller.GetClinicsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message string                   `json:"message"`
				Responses   []dto.ClinicResponse 	 `json:"response"`
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