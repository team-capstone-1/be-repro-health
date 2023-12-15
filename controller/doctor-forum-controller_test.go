package controller_test

import (
	"capstone-project/controller"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func TestGetDoctorAllForumsController(t *testing.T) {
// 	// Setup Echo instance
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/forums", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Set up your database for testing
// 	database.InitTest()

// 	// Call the controller function
// 	err := controller.GetDoctorAllForumsController(c)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	// Add more assertions as needed
// }

func TestFailedGetDoctorForumDetails(t *testing.T) {
	e := echo.New()

	forumID := uuid.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = "doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/forums/details/"+forumID.String(), nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/forums/details/" + forumID.String())

	c.Set("doctor", token)

	err = controller.GetDoctorForumDetails(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateDoctorReplyForumUnauthorized(t *testing.T) {
	e := echo.New()

	// Create a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "doctor_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Prepare the request body
	requestBody := `{
        "forum_id": "3b1d0a0f-7b0b-4c0d-8e1e-2e1e1e1e1e1e",
    	"content": "Semoga jawaban saya membantu Anda"
    }`

	req := httptest.NewRequest(http.MethodPost, "/doctors/forum-replies", strings.NewReader(requestBody))
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set the JWT token directly in the context
	c.Set("user", token)

	// Execute the correct controller for creating a doctor reply to a forum
	err := controller.CreateDoctorReplyForum(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	// Additional assertions based on the response if needed
}

// func TestCreateDoctorReplyForum(t *testing.T) {
// 	e := echo.New()

// 	// ...
// 	jwtKey := os.Getenv("JWT_KEY")
// 	doctorID := uuid.MustParse("f7613c10-29fd-4b82-bfea-1649ae41af98")
// 	// Create a JWT token
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["user_id"] = doctorID
// 	claims["name"] = "Rizki"
// 	claims["role"] = "doctor"
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

// 	tokenString, err := token.SignedString([]byte(jwtKey))
// 	if err != nil {
// 		t.Fatalf("Error creating JWT token: %v", err)
// 	}

// 	// Prepare the request body
// 	requestObject := dto.DoctorForumReplyRequest{
// 		ForumsID: uuid.MustParse("3b1d0a0f-7b0b-4c0d-8e1e-2e1e1e1e1e1e"),
// 		DoctorID: doctorID,
// 		Content:  "Test",
// 	}

// 	// fmt.Printf("Token: %s\n", tokenString)

// 	requestBody, err := json.Marshal(requestObject)
// 	if err != nil {
// 		// Handle error jika terjadi kesalahan dalam marshalling
// 		panic(err)
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/doctors/forum-replies", strings.NewReader(string(requestBody)))
// 	req.Header.Set("Authorization", "Bearer "+tokenString)
// 	req.Header.Set("Content-Type", "application/json")

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Set the JWT token directly in the context
// 	c.Set("user", tokenString)

// 	// Set the path and method for the context
// 	c.SetPath("/doctors/forum-replies")

// 	// Apply CheckRole middleware
// 	role := "doctor"
// 	checkRoleMiddleware := middleware.CheckRole(role)
// 	handler := checkRoleMiddleware(controller.CreateDoctorReplyForum)

// 	// Execute the correct controller for creating a doctor reply to a forum
// 	err = handler(c)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusCreated, rec.Code)
// 	// Tambahkan setelah pengujian assertions
// 	// fmt.Println("RESPONSE BODY:", rec.Body.String())

// }
