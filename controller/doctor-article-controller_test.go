package controller_test

import (
	"bytes"
	"capstone-project/config"
	"capstone-project/controller"
	"capstone-project/dto"
	"capstone-project/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
		article    model.Article
		expertCode int
	}{
		{
			name: "create new article",
			path: "/articles",
			article: model.Article{
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
	token := InsertDataArticle

	for _, testCase := range testCases {
		articleJson, _ := json.Marshal(testCase.article)

		req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(string(articleJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// req.Header.Set(echo.HeaderAuthorization, "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJuYW1lIjoiRGF2aW5ubiIsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6IjUwMWM3MzdhLTcyY2EtNGY1ZS04YjM1LWY1Mzc0ZTRmZDg1YyJ9.Ioa0l1n0vJpqi0BrQOWT0skSEMMGxi49g_y3_QrBh0w")
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(config.JWT_KEY))(controller.CreateDoctorArticleControllerTesting())(context)
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, controller.CreateDoctorArticleController(c)) {
			assert.Equal(t, testCase.expertCode, rec.Code)
			body := rec.Body.String()

			type Response struct {
				Message  string                   `json:"message"`
				Response dto.DoctorArticleRequest `json:"response"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)
			fmt.Println("token:", responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}

			if rec.Code == 200 {
				assert.Equal(t, testCase.article.Title, responseData.Response.Title)
				assert.Equal(t, testCase.article.Tags, responseData.Response.Tags)
				assert.Equal(t, testCase.article.Reference, responseData.Response.Reference)
				assert.Equal(t, testCase.article.Image, responseData.Response.Image)
				assert.Equal(t, testCase.article.ImageDesc, responseData.Response.ImageDesc)
				assert.Equal(t, testCase.article.Content, responseData.Response.Content)
			}
		}
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
