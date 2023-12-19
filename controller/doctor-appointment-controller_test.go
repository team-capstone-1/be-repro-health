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

func TestDoctorGetAllConsultations(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get all appointments",
		// 	path:       "/doctors/appointments/details-consultation",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get all appointments",
			path:       "/doctors/appointments/details-consultations",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	token, doctor := InsertDataDoctor()
	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, "/doctors/appointments/details-consultation", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorGetAllConsultationsTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDoctorGetDetailsPatientController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get details consultation",
		// 	path:       "/doctors/appointments/details-consultation/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get details consultation",
			path:       "/doctors/appointmentss/details-consultation/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, doctor := InsertDataDoctor()
	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, "/doctors/appointments/details-consultation/:id", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorGetDetailsPatientControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /doctors/appointments/details-consultation/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDoctorGetDetailsTransactionController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get details transactions",
		// 	path:       "/doctors/appointments/details-transaction/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get details transactions",
			path:       "/doctors/appointments/details-transactions/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, doctor := InsertDataDoctor()
	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, "/doctors/appointments/details-transaction/:id", nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorGetDetailsTransactionControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("GET /doctors/appointments/details-transaction/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
func TestDoctorConfirmConsultationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		patient    dto.DoctorConfirmConsultationRequest
		expectCode int
	}{
		// {
		// 	name: "success confirm consultation",
		// 	path: "/doctors/appointments/confirm-consultation",
		// 	patient: dto.DoctorConfirmConsultationRequest{
		// 		ConsultationID: uuid.New(),
		// 	},
		// 	expectCode: http.StatusOK,
		// },
		{
			name: "failed confirm consultation",
			path: "/doctors/appointments/confirm-consultations",
			patient: dto.DoctorConfirmConsultationRequest{
				ConsultationID: uuid.New(),
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()
	InsertDataUser()
	InsertDataPatient(uuid.New())

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPut, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		// context.SetParamNames("id")
		// context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorConfirmConsultationControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("PUT /doctors/appointments/confirm-consultation", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDoctorFinishedConsultationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		patient    dto.DoctorConfirmConsultationRequest
		expectCode int
	}{
		// {
		// 	name: "success finish consultation",
		// 	path: "/doctors/appointments/finish-consultation",
		// 	patient: dto.DoctorConfirmConsultationRequest{
		// 		ConsultationID: uuid.New(),
		// 	},
		// 	expectCode: http.StatusOK,
		// },
		{
			name: "failed finish consultation",
			path: "/doctors/appointments/finish-consultation",
			patient: dto.DoctorConfirmConsultationRequest{
				ConsultationID: uuid.New(),
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()
	InsertDataUser()
	InsertDataPatient(uuid.New())

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodPut, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set("Content-Type", writer.FormDataContentType())

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		// context.SetParamNames("id")
		// context.SetParamValues(doctor.ID.String())
		middleware.JWT([]byte(config.JWT_KEY))(controller.DoctorFinishedConsultationControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("PUT /doctors/appointments/finish-consultation", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
