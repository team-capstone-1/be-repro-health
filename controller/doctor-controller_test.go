package controller_test

import (
	"capstone-project/config"
	"capstone-project/constant"
	"capstone-project/controller"
	"capstone-project/database"
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func InsertDataAdmin() string {
	user := model.User{
		ID:       uuid.New(),
		Name:     "Admin1",
		Email:    "admin1@gmail.com",
		Password: "Admin@123",
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	database.DB.Create(&user)

	token, _ := m.CreateToken(user.ID, constant.ROLE_ADMIN, user.Name, false)
	return token
}

func InsertDataClinics() {
	clinicID, _ := uuid.Parse("fd9d3738-3bcf-4693-9079-57ad3b375af5")
	// CLINIC SEEDERS
	clinic := []model.Clinic{
		{
			ID:       clinicID,
			Name:     "Klinik Inter Medika",
			City:     "Jakarta Pusat",
			Location: "Jl. K.S. Tubun No.21, Petamburan, Kota Jakarta Pusat ",
			Profile:  "Selamat datang di Klinik Inter Medika, tempat di mana kesehatan dan kenyamanan pasien menjadi prioritas utama. Klinik kami berkomitmen untuk memberikan pelayanan medis berkualitas tinggi.",
		},
	}

	for _, v := range clinic {
		var exist model.Clinic

		errCheck := database.DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			database.DB.Create(&v)
		}
	}
	return
}

func InsertDataSpecialist() {
	specialistID, _ := uuid.Parse("0a8f12c6-1244-43e6-a704-126b173a8732")

	// SPECIALIST SEEDERS
	specialist := []model.Specialist{
		{
			ID:    specialistID,
			Name:  "Kandungan",
			Image: "https://res.cloudinary.com/dw3n2ondc/image/upload/v1700464911/Reproduction-Health/rrau8asadj80uuugksdj.png",
		},
	}

	for _, v := range specialist {
		var exist model.Specialist

		errCheck := database.DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			database.DB.Create(&v)
		}
	}
	return
}

func InsertDataDoctor() string {
	// user := m.ExtractTokenUserId(c)
	specialistID, _ := uuid.Parse("0a8f12c6-1244-43e6-a704-126b173a8732")
	clinicID, _ := uuid.Parse("fd9d3738-3bcf-4693-9079-57ad3b375af5")
	doctorID, _ := uuid.Parse("f7613c10-29fd-4b82-bfea-1649ae41af98")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("Doctor@123"), bcrypt.DefaultCost)
	doctor := model.Doctor{

		ID:           doctorID,
		Name:         "Dr. Andi Cahaya, Sp.OG",
		Email:        "andicahyo@gmail.com",
		Password:     string(passwordHash),
		Price:        150000,
		Address:      "Klinik Nasional. Jl. Bedrek No.47e, Sanggrahan, Condongcatur, Kec. Depok, Kabupaten Sleman, DIY",
		Phone:        "+62 812345865",
		ProfileImage: "https://res.cloudinary.com/dw3n2ondc/image/upload/v1700464491/Reproduction-Health/kypalfa66lmowdu5zh9p.png",
		SpecialistID: specialistID,
		ClinicID:     clinicID,
	}

	InsertDataClinics()
	InsertDataSpecialist()

	database.DB.Create(&doctor)

	token, _ := m.CreateToken(doctor.ID, constant.ROLE_DOCTOR, doctor.Name, false)
	return token
}

func TestCreateDoctorController(t *testing.T) {

	specialistID, _ := uuid.Parse("0a8f12c6-1244-43e6-a704-126b173a8732")
	clinicID, _ := uuid.Parse("fd9d3738-3bcf-4693-9079-57ad3b375af5")
	var testCases = []struct {
		name       string
		path       string
		doctor     dto.DoctorSignUpRequest
		expectCode int
	}{
		// {
		// 	name: "create new doctor",
		// 	path: "/admins/doctors/signup",
		// 	doctor: dto.DoctorSignUpRequest{
		// 		Name:         "Dr. Thamrin Hakamoto, Sp.Bedah",
		// 		Email:        "hakamoto@gmail.com",
		// 		Password:     "haka@123",
		// 		Price:        150000,
		// 		Address:      "Klinik Bougenville. Jl. Asia Afrika No.113, Gegerkalong, Parahyangan, Kec. Antasari, Kota Bandung, Jawa Barat",
		// 		Phone:        "+62 877694502343",
		// 		ProfileImage: "",
		// 		SpecialistID: specialistID,
		// 		ClinicID:     clinicID,
		// 	},
		// 	expectCode: http.StatusCreated,
		// },
		// {
		// 	name: "failed create new doctor invalid endpoint",
		// 	path: "/admins/doctors/signups",
		// 	doctor: dto.DoctorSignUpRequest{
		// 		Name:         "Dr. Thamrin Hakamoto, Sp.Bedah",
		// 		Email:        "hakamoto@gmail.com",
		// 		Password:     "haka@123",
		// 		Price:        150000,
		// 		Address:      "Klinik Bougenville. Jl. Asia Afrika No.113, Gegerkalong, Parahyangan, Kec. Antasari, Kota Bandung, Jawa Barat",
		// 		Phone:        "+62 877694502343",
		// 		ProfileImage: "",
		// 		SpecialistID: specialistID,
		// 		ClinicID:     clinicID,
		// 	},
		// 	expectCode: http.StatusNotFound,
		// },
		{
			name: "failed create new doctor email existed",
			path: "/admins/doctors/signup",
			doctor: dto.DoctorSignUpRequest{
				Name:         "Dr. Thamrin Hakamoto, Sp.Bedah",
				Email:        "hakamoto@gmail.com",
				Password:     "haka@123",
				Price:        150000,
				Address:      "Klinik Bougenville. Jl. Asia Afrika No.113, Gegerkalong, Parahyangan, Kec. Antasari, Kota Bandung, Jawa Barat",
				Phone:        "+62 877694502343",
				ProfileImage: "",
				SpecialistID: specialistID,
				ClinicID:     clinicID,
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataClinics()
	InsertDataSpecialist()
	token := InsertDataAdmin()

	fmt.Print(token)

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.doctor)

		req := httptest.NewRequest(http.MethodPost, "/admins/doctors/signup", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.SignUpDoctorControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run("POST /admins/doctors/signup", func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}

}

func TestDoctorLoginController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user       model.User
		expectCode int
	}{
		{
			name: "Doctor Login",
			path: "/doctors/login",
			user: model.User{
				Email:    "andicahyo@gmail.com",
				Password: "Doctor@123",
			},
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)

		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.DoctorLoginController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type LoginResponse struct {
				UserID string `json:"user_id"`
				Email  string `json:"email"`
				Name   string `json:"name"`
				Token  string `json:"token"`
			}
			type Response struct {
				Message  string        `json:"message"`
				Response LoginResponse `json:"response"`
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

func TestDoctorSendOTPController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user       model.User
		expectCode int
	}{
		{
			name: "Doctor Success Send OTP",
			path: "/doctors/send-otp",
			user: model.User{
				Email: "andicahyo@gmail.com",
			},
			expectCode: http.StatusOK,
		},
		{
			name: "Doctor Failed Send OTP",
			path: "/doctors/send-otp",
			user: model.User{
				Email: "andicahyo1@gmail.com",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)

		req := httptest.NewRequest(http.MethodPut, testCase.path, strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.DoctorSendOTPController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message  string         `json:"message"`
				Response map[string]any `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestDoctorValidateOTPController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user       model.User
		expectCode int
	}{
		{
			name: "Doctor Validate OTP",
			path: "/doctors/validate-otp",
			user: model.User{
				Email: "andicahyo@gmail.com",
				OTP:   "",
			},
			expectCode: http.StatusOK,
		},
		{
			name: "Doctor Validate OTP Invalid OTP",
			path: "/doctors/validate-otp",
			user: model.User{
				Email: "andicahyo@gmail.com",
				OTP:   "1",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)

		req := httptest.NewRequest(http.MethodPut, testCase.path, strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.DoctorValidateOTPController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message  string         `json:"message"`
				Response map[string]any `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestGetDoctorByIDController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get details doctor",
		// 	path:       "/doctors/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "success get details doctor",
			path:       "/doctorss/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path+"f7613c10-29fd-4b82-bfea-1649ae41af98", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.GetDoctorController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message  string           `json:"message"`
				Response dto.UserResponse `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}

		}
	}
}

func TestGetDoctorsBySpecialistController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get specialist",
		// 	path:       "/specialists/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get specialist",
			path:       "/specialists/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()
	InsertDataSpecialist()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.GetDoctorsBySpecialistController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message  string           `json:"message"`
				Response dto.UserResponse `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}

		}
	}
}

func TestGetDoctorsByClinicController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get clinics",
		// 	path:       "/clinics/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get clinics",
			path:       "/clinics/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()
	InsertDataClinics()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.GetDoctorsByClinicController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message  string           `json:"message"`
				Response dto.UserResponse `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}

		}
	}
}

func TestGetDoctorsBySpecialistAndClinicController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		// {
		// 	name:       "success get doctor by clinics and specialist",
		// 	path:       "/clinics/:id/specialists/:id",
		// 	expectCode: http.StatusOK,
		// },
		{
			name:       "failed get doctor by clinics and specialist",
			path:       "/clinics/:id/specialists/:id",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataDoctor()
	InsertDataClinics()
	InsertDataSpecialist()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.GetDoctorsBySpecialistAndClinicController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message  string           `json:"message"`
				Response dto.UserResponse `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}

		}
	}
}

func TestChangeDoctorPasswordController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		doctor     dto.ChangeDoctorPasswordRequest
		expectCode int
	}{
		{
			name: "success update password",
			path: "/doctors/change-password",
			doctor: dto.ChangeDoctorPasswordRequest{
				Password: "Rizki@123",
			},
			expectCode: http.StatusBadRequest,
		},
		{
			name: "failed update password",
			path: "/doctors/change-passwords",
			doctor: dto.ChangeDoctorPasswordRequest{
				Password: "Rizki@123",
			},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	token := InsertDataDoctor()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.doctor)

		req := httptest.NewRequest(http.MethodPut, testCase.path, strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.ChangeDoctorPasswordControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("POST %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}

}
