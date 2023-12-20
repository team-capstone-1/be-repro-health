package controller_test

import (
	"capstone-project/controller"
	"capstone-project/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAdminLoginController(t *testing.T) {
	// Set up a new Echo instance
	e := echo.New()

	// Prepare a sample login request
	loginReq := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Create a request
	req := httptest.NewRequest(http.MethodPost, "/admin/login", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Bind the login request to the context
	ctx.SetPath("/admin/login")
	assert.NoError(t, ctx.Bind(&loginReq))

	// Call the controller function
	err := controller.AdminLoginController(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
