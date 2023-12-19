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

func TestCreateConsultationController(t *testing.T) {
	e := InitEchoTestAPI()
	token, user := InsertDataUser()
	patient, _ := InsertDataPatient(user.ID)
	var testCases = []struct {
		name       string
		path       string
		patient    dto.ConsultationRequest
		expectCode int
	}{
		// {
		// 	name: "success create new consultation",
		// 	path: "/consultations",
		// 	patient: dto.ConsultationRequest{
		// 		PatientID:     patient.ID,
		// 		DoctorID:      uuid.New(),
		// 		Date:          "2006-06-02",
		// 		Session:       "siang",
		// 		PaymentMethod: "manual_transfer",
		// 	},
		// 	expectCode: http.StatusCreated,
		// },
		{
			name: "failed create new consultation",
			path: "/consultationss",
			patient: dto.ConsultationRequest{
				PatientID:     patient.ID,
				DoctorID:      uuid.New(),
				Date:          "2006-06-02",
				Session:       "siang",
				PaymentMethod: "manual_transfer",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(patient.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateConsultationControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("POST /consultations", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
