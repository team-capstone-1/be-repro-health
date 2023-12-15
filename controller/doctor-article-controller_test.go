package controller_test

import (
	"capstone-project/controller"
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

func TestGetAllArticleDoctorsController_Unauthorized(t *testing.T) {
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

	fmt.Println("Response Body:", rec.Body.String())
}

func TestGetDoctorArticleByIDController(t *testing.T) {
	e := echo.New()

	articleID := uuid.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "doctor_id"
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

// func TestCreateDoctorArticleController(t *testing.T) {
// 	e := echo.New()

// 	// Set up test JWT key
// 	jwtKey := os.Getenv("JWT_KEY")

// 	// Create a JWT token with a valid doctor ID
// 	token := createTestToken(jwtKey)

// 	// Create a request body for the article creation
// 	articleRequest := dto.DoctorArticleRequest{
// 		Title:       "Test Article",
// 		Description: "This is a test article.",
// 	}

// 	// Convert the article request to JSON
// 	articleJSON, err := util.ConvertToJSON(articleRequest)
// 	assert.NoError(t, err)

// 	// Create a multipart/form-data request with the token and article JSON
// 	req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(articleJSON))
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "multipart/form-data")

// 	// Create an HTTP recorder and an Echo context
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Call the controller function
// 	err = controller.CreateDoctorArticleController(c)

// 	// Assert that there is no error (indicating success) and the response code is StatusCreated
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusCreated, rec.Code)

// 	// Optionally, you can inspect the response body for more details
// 	fmt.Println("Response Body:", rec.Body.String())
// }

// func createTestToken(jwtKey string) string {
// 	token := util.CreateTestToken(jwtKey, "valid_doctor_id")
// 	tokenString, err := token.SignedString([]byte(jwtKey))
// 	if err != nil {
// 		panic(fmt.Sprintf("Error creating test JWT token: %v", err))
// 	}
// 	return tokenString
// }
