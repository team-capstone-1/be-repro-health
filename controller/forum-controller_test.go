package controller_test

import (
	"capstone-project/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetForumsController(t *testing.T) {
	// Initialize Echo
	e := echo.New()

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
	assert.Equal(t, http.StatusOK, rec.Code)

	// Add more assertions based on your specific logic and expected outcomes
}
