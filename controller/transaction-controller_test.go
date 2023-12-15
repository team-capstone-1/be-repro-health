package controller_test

import (
	"capstone-project/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/transactions/123", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/transactions/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")
	assert.NoError(t, controller.GetTransactionController(c))

	err := controller.GetTransactionController(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// func TestGetTransactionsController(t *testing.T) {
// 	e := echo.New()

// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)
// 	c.SetPath("")
// 	c.SetParamNames("transactions")
// 	c.SetParamValues("")

// 	// Call the controller function
// 	err := controller.GetTransactionsController(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusBadRequest, rec.Code)
// }