package controller_test

import (
	"bytes"
	"capstone-project/controller"
	"capstone-project/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionController_invalid(t *testing.T) {
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

func TestGetTransactionsController_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "user_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/transactions"+tokenString, nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions" + tokenString)

	c.Set("user", token)

	err = controller.GetTransactionsController(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestGetPatientTransactionsController_invalid(t *testing.T) {
	e := echo.New()

	jwtKey := os.Getenv("JWT_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "user_id"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		t.Fatalf("Error creating JWT token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/transactions"+tokenString, nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions" + tokenString)

	c.Set("user", token)

	err = controller.GetPatientTransactionsController(c)

	assert.Nil(t, err, "Expected an error but got nil")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreatePaymentController_invalid(t *testing.T) {
	e := echo.New()

	transactionID := uuid.New()

	paymentRequest := dto.PaymentRequest{
		Name:          "John Doe",
		AccountNumber: "123456789",
		Image:         "base64_encoded_image_data",
	}

	paymentJSON, err := json.Marshal(paymentRequest)
	assert.NoError(t, err, "Error marshalling payment request to JSON")

	req := httptest.NewRequest(http.MethodPost, "/transactions/"+transactionID.String(), bytes.NewBuffer(paymentJSON))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions" + transactionID.String())

	err = controller.CreatePaymentController(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response body")
	assert.Equal(t, "error parse id", response["message"])
}
