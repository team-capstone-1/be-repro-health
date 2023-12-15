package controller_test

// import (
// 	"capstone-project/controller"
// 	"capstone-project/database"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

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

// func TestGetDoctorForumDetails(t *testing.T) {
// 	// Setup Echo instance
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/path", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/forums/details/:id")
// 	c.SetParamNames("id")
// 	c.SetParamValues("your_forum_id_here")

// 	// Set up your database for testing
// 	database.InitTest()

// 	// Call the controller function
// 	err := controller.GetDoctorForumDetails(c)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	// Add more assertions as needed
// }

// func TestCreateDoctorReplyForum(t *testing.T) {
// 	// Setup Echo instance
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(`{"forum_id": "your_forum_id_here", "doctor_id": "your_doctor_id_here", "content": "your_content_here", "date": "your_date_here"}`))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Set up your database for testing
// 	database.InitTest()

// 	// Call the controller function
// 	err := controller.CreateDoctorReplyForum(c)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	// Add more assertions as needed
// }
