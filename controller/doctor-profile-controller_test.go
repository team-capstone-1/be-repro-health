package controller_test

import (
	"bytes"
	"capstone-project/config"
	"capstone-project/constant"
	"capstone-project/controller"
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		panic("Error loading .env file")
	}
}

func TestGetDoctorProfileController(t *testing.T) {}

func InsertDataDoctorWorkHistory() (string, error) {
	id := uuid.New()
	doctorID := uuid.New()

	startingDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endingDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	doctorWorkHistory := model.DoctorWorkHistory{
		ID:           id,
		DoctorID:     doctorID,
		StartingDate: startingDate,
		EndingDate:   endingDate,
		Job:          "Job",
		Workplace:    "Workplace",
		Position:     "Position",
	}

	token, err := m.CreateToken(doctorWorkHistory.DoctorID, constant.ROLE_DOCTOR, "DoctorName", false)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Work History
func TestGetDoctorWorkHistoriesController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get doctor work history",
			path:       "/doctors/work-history",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorWorkHistory()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/doctors/work-history", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDoctorWorkHistoriesControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestCreateDoctorWorkHistoryController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		dwh        dto.DoctorWorkHistoryRequest
		expectCode int
	}{
		{
			name: "success create doctor work history",
			path: "/doctors/work-history",
			dwh: dto.DoctorWorkHistoryRequest{
				DoctorID:     uuid.New(),
				StartingDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				Job:          "Job",
				Workplace:    "Workplace",
				Position:     "Position",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "test_image_url.jpg")
		if err != nil {
			t.Fatalf(err.Error())
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatalf(err.Error())
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/doctors/work-history", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateDoctorWorkHistoryControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("POST /doctors/work-history", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestUpdateDoctorWorkHistoryController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		dwh        dto.DoctorWorkHistoryRequest
		expectCode int
	}{
		{
			name: "update work history",
			path: "/doctors/work-history/1",
			dwh: dto.DoctorWorkHistoryRequest{
				DoctorID:     uuid.New(),
				StartingDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				Job:          "Job in new workplace",
				Workplace:    "Workplace in new workplace",
				Position:     "Position in new workplace",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorWorkHistory()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "test_image_url.jpg")
		if err != nil {
			t.Fatalf(err.Error())
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatalf(err.Error())
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/doctors/work-history/:id", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdateDoctorArticleControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("PUT /doctors/work-history/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeleteDoctorWorkHistoryController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete work history by id",
			path:       "/doctors/work-history/1", // Replace with a valid work history ID
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorWorkHistory()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, testCase.path, nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath("/doctors/work-history/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.DeleteDoctorWorkHistoryControllerTesting())(context)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func InsertDataDoctorEducation() (string, error) {
	id := uuid.New()
	doctorID := uuid.New()

	startingDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endingDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	doctorEducation := model.DoctorEducation{
		ID:               id,
		DoctorID:         doctorID,
		StartingDate:     startingDate,
		EndingDate:       endingDate,
		EducationProgram: "Education Program",
		University:       "University",
	}

	token, err := m.CreateToken(doctorEducation.DoctorID, constant.ROLE_DOCTOR, "DoctorName", false)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Education
func TestGetDoctorEducationsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get doctor educations",
			path:       "/doctors/educations",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorEducation()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/doctors/educations", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDoctorEducationsControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestCreateDoctorEducationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		dwh        dto.DoctorEducationRequest
		expectCode int
	}{
		{
			name: "success create doctor educations",
			path: "/doctors/educations",
			dwh: dto.DoctorEducationRequest{
				DoctorID:         uuid.New(),
				StartingDate:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				EducationProgram: "Education Program",
				University:       "University",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "test_image_url.jpg")
		if err != nil {
			t.Fatalf(err.Error())
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatalf(err.Error())
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/doctors/educations", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateDoctorEducationsControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("POST /doctors/educations", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestUpdateDoctorEducationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		dwh        dto.DoctorEducationRequest
		expectCode int
	}{
		{
			name: "update work history",
			path: "/doctors/work-history/1",
			dwh: dto.DoctorEducationRequest{
				DoctorID:         uuid.New(),
				StartingDate:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingDate:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				EducationProgram: "Education Program",
				University:       "University",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorEducation()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "test_image_url.jpg")
		if err != nil {
			t.Fatalf(err.Error())
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatalf(err.Error())
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/doctors/education/:id", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdateDoctorEducationControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("PUT /doctors/education/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeleteDoctorEducationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete education by id",
			path:       "/doctors/education/1", // Replace with a valid work history ID
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorEducation()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, testCase.path, nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath("/doctors/education/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.DeleteDoctorEducationControllerTesting())(context)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

// Certification
func InsertDataDoctorCertification() (string, error) {
	id := uuid.New()
	doctorID := uuid.New()

	startingDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endingDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	doctorCertification := model.DoctorCertification{
		ID:              id,
		DoctorID:        doctorID,
		CertificateType: "Certificate Type",
		Description:     "Description",
		StartingDate:    startingDate,
		EndingDate:      endingDate,
		FileSize:        "FileSize",
		Details:         "Details",
	}

	token, err := m.CreateToken(doctorCertification.DoctorID, constant.ROLE_DOCTOR, "DoctorName", false)
	if err != nil {
		return "", err
	}

	return token, nil
}

func TestGetDoctorCertificationsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get doctor certifications",
			path:       "/doctors/certifications",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorCertification()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/doctors/certifications", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDoctorCertificationsControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestCreateDoctorCertificationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		dwh        dto.DoctorCertificationRequest
		expectCode int
	}{
		{
			name: "success create doctor certification",
			path: "/doctors/certification",
			dwh: dto.DoctorCertificationRequest{
				DoctorID:        uuid.New(),
				StartingDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				Description:     "Description",
				CertificateType: "Certificate Type",
				Details:         "Details",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "test_image_url.jpg")
		if err != nil {
			t.Fatalf(err.Error())
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatalf(err.Error())
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/doctors/certification", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateDoctorCertificationControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("POST /doctors/educations", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestUpdateDoctorCertificationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		dwh        dto.DoctorCertificationRequest
		expectCode int
	}{
		{
			name: "update doctor certification",
			path: "/doctors/certification/1",
			dwh: dto.DoctorCertificationRequest{
				DoctorID:        uuid.New(),
				StartingDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndingDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				Description:     "Description updated",
				CertificateType: "Certificate Type updated",
				Details:         "Details updated",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorCertification()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		fileWriter, err := writer.CreateFormFile("image", "test_image_url.jpg")
		if err != nil {
			t.Fatalf(err.Error())
		}

		file, err := os.Open("../docs/ERD.png")
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			t.Fatalf(err.Error())
		}

		writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/doctors/certification/:id", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdateDoctorCertificationControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("PUT /doctors/certification/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestDeleteDoctorCertificationController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete doctor certification by id",
			path:       "/doctors/certification/1", // Replace with a valid work history ID
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctorCertification()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, testCase.path, nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath("/doctors/certification/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.DeleteDoctorCertificationControllerTesting())(context)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}
