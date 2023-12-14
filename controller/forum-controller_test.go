package controller_test

import (
	"capstone-project/controller"
	"capstone-project/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// / Mock a function to generate a valid token for testing
func generateValidToken(userID uuid.UUID) string {
	token, _ := middleware.CreateToken(userID, "user", "John Doe", true)
	return token
}
func TestGetForumsController(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Inisialisasi objek controller
	// controller := &ForumController{}

	req := httptest.NewRequest(http.MethodGet, "/forums", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/forums")

	// Call the controller function
	err := controller.GetForumsController(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetForumController(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Create a new request with a valid forum ID
	forumID := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/forums/"+forumID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the controller function
	err := controller.GetForumController(c)

	c.SetPath("/forums" + forumID.String())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateForumController(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Mock a valid user ID for the token extraction
	validUserID := uuid.New()
	token := generateValidToken(validUserID)

	// Create a new request with valid forum data
	req := httptest.NewRequest(http.MethodPost, "/forums", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Set("Authorization", "Bearer "+token)

	// Mock the JWT middleware with a valid token
	middlewareMock := middleware.CheckRole("user")
	handler := func(c echo.Context) error {
		return controller.CreateForumController(c)
	}

	// Call the middleware and controller function
	err := middlewareMock(handler)(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestDeleteForumController(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Mock a valid user ID for the token extraction
	validUserID := uuid.New()
	token := generateValidToken(validUserID)

	// Mock a valid forum ID
	forumID := uuid.New()

	// Create a new request with valid forum ID
	req := httptest.NewRequest(http.MethodDelete, "/forums/"+forumID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Set("Authorization", "Bearer "+token)

	// Call the controller function
	err := controller.DeleteForumController(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
