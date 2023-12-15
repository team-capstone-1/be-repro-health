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

func TestGetTransactionController(t *testing.T) {
	e := echo.New()

	transactionID := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/transactions/"+transactionID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.GetTransactionController(c)

	c.SetPath("/transactions" + transactionID.String())

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
