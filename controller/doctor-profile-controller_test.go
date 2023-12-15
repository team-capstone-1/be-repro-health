// controller_test.go
package controller_test

// import (
// 	"capstone-project/controller"
// 	m "capstone-project/middleware"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/google/uuid"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func ExtractTokenUserId(e echo.Context) uuid.UUID {
// 	// Your implementation here
// 	return uuid.New()
// }

// func TestGetDoctorProfileController(t *testing.T) {
// 	// Set up a new Echo instance
// 	e := echo.New()

// 	// Create a request
// 	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
// 	rec := httptest.NewRecorder()
// 	ctx := e.NewContext(req, rec)

// 	// Mock ExtractTokenUserId middleware
// 	token := &jwt.Token{} // This is a mock JWT token
// 	m.ExtractTokenUserId = func(c echo.Context) (string, error) {
// 		// You might want to set some claims or information in the token for your tests
// 		c.Set("user", token)
// 		return "user123", nil
// 	}

// 	// Call the controller function
// 	err := controller.GetDoctorProfileController(ctx)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)

// 	// Convert the expected response map to a JSON string
// 	expectedResponse := `{
// 		"message": "success get doctor profile",
// 		"response": {
// 			"id": "doctor123",
// 			"name": "Dr. John Doe",
// 			"specialty": "Cardiologist",
// 			"description": "Experienced cardiologist with a focus on heart health."
// 		}
// 	}`

// 	assert.JSONEq(t, expectedResponse, rec.Body.String())
// }
