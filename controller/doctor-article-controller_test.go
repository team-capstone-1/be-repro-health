package controller_test

import (
	"bytes"
	"capstone-project/controller"
	"capstone-project/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
