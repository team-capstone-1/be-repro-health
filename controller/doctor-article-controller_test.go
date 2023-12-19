package controller_test

import (
	"bytes"
	"capstone-project/config"
	"capstone-project/constant"
	"capstone-project/controller"
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
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

func InsertDataArticle() (string, error) {
	// Create a new Article instance
	article := model.Article{
		ID:        uuid.New(),
		DoctorID:  uuid.New(),
		Title:     "Test Article",
		Tags:      "tag1, tag2",
		Reference: "Reference",
		Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ...",
		Image:     "image_url.jpg",
		ImageDesc: "Image Description",
		Date:      time.Now(),
		Published: true,
		View:      0,
		Comment:   []model.Comment{}, // Adjust based on your Comment model.
	}

	// Create a token for the doctor
	token, err := m.CreateToken(article.DoctorID, constant.ROLE_DOCTOR, "DoctorName", false)
	if err != nil {
		return "", err
	}

	return token, nil
}

func TestGetAllArticleDoctorsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get all articles",
			path:       "/articles",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataArticle()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetAllArticleDoctorsControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetAllArticleDoctorsController_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", token)

	err = controller.GetAllArticleDoctorsController(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestGetDoctorArticleByIDController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get article by id",
			path:       "/articles/1", // Replace with a valid article ID
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataArticle()
	token, _ := InsertDataDoctor()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		req.Header.Set("Authorization", "Bearer "+token)

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath("/articles/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetDoctorArticleByIDControllerTesting())(context)

		t.Run(testCase.name, func(t *testing.T) {
			fmt.Printf("Actual status code: %d\n", rec.Code) // Debugging statement
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestGetDoctorArticleByIDController_invalid(t *testing.T) {
	e := echo.New()

	articleID := uuid.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/articles/"+articleID.String(), nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles" + articleID.String())

	c.Set("user", token)

	err = controller.GetDoctorArticleByIDController(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestCreateDoctorArticleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		article    dto.DoctorArticleRequest
		expertCode int
	}{
		{
			name: "create new article",
			path: "/articles",
			article: dto.DoctorArticleRequest{
				Title:     "Test Article",
				Tags:      "Test, Article, Tag",
				Reference: "Test Reference",
				Image:     "test_image_url.jpg",
				ImageDesc: "Test Image Description",
				Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			},
			expertCode: http.StatusCreated,
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

		req := httptest.NewRequest(http.MethodPost, "/articles", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateDoctorArticleControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run("POST /articles", func(t *testing.T) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})
	}
}

func TestCreateDoctorArticleController_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	articleDTO := dto.DoctorArticleRequest{
		Title:   "Test Article",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	jsonArticleDTO, err := json.Marshal(articleDTO)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(jsonArticleDTO))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("user", token)

	err = controller.CreateDoctorArticleController(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestUpdateDoctorArticleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		article    dto.DoctorArticleRequest
		expertCode int
	}{
		{
			name: "update article",
			path: "/articles/1",
			article: dto.DoctorArticleRequest{
				Title:   "Updated Test Article",
				Content: "Updated content goes here.",
			},
			expertCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataArticle()
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

		req := httptest.NewRequest(http.MethodPut, "/articles/:id", &buf)
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

		t.Run("PUT /articles/:id", func(t *testing.T) {
			assert.Equal(t, testCase.expertCode, rec.Code)
		})
	}
}

func TestUpdateDoctorArticleController_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	articleID := "some_article_id"
	updateArticleDTO := dto.DoctorArticleRequest{
		Title:   "Updated Test Article",
		Content: "Updated content goes here.",
	}

	jsonUpdateArticleDTO, err := json.Marshal(updateArticleDTO)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/articles/"+articleID, bytes.NewBuffer(jsonUpdateArticleDTO))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles/" + articleID)

	c.Set("user", token)

	err = controller.UpdateDoctorArticleController(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateArticlePublishedStatusController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		article    dto.DoctorArticleResponse
		expertCode int
	}{
		{
			name: "update article",
			path: "/articles/1/publish",
			article: dto.DoctorArticleResponse{
				Published: true,
			},
			expertCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataArticle()
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

		req := httptest.NewRequest(http.MethodPut, "/articles/:id/publish", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))

		rec := httptest.NewRecorder()

		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")

		middleware.JWT([]byte(config.JWT_KEY))(controller.UpdateArticlePublishedStatusControllerTesting())(context)
		t.Run("PUT /articles/:id/publish", func(t *testing.T) {
			assert.Equal(t, testCase.expertCode, rec.Code)
		})
	}
}

func TestUpdateArticlePublishedStatusController_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "invalid_doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	articleID := "some_article_id"
	updateStatusDTO := dto.DoctorArticleResponse{
		Published: true,
	}

	jsonUpdateStatusDTO, err := json.Marshal(updateStatusDTO)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/articles/"+articleID+"/publish", bytes.NewBuffer(jsonUpdateStatusDTO))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles/" + articleID + "/publish")

	c.Set("user", token)

	err = controller.UpdateArticlePublishedStatusController(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
