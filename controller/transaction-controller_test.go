package controller_test

import (
	"bytes"
	"capstone-project/config"
	"capstone-project/constant"
	"capstone-project/controller"
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InsertDataTransaction() (string, error) {
	transaction := model.Transaction{
		ID:             uuid.New(),
		ConsultationID: uuid.New(),
		Invoice:        "Test Invoice",
		Status:         model.Waiting,
		PaymentStatus:  "pending",
		Price:          0.0,
		AdminPrice:     0.0,
		Total:          0.0,
	}

	// Insert the new transaction into the database
	token, err := m.CreateToken(transaction.ConsultationID, constant.ROLE_USER, "UserName", false)
	if err != nil {
		return "", err
	}

	return token, nil
}

func TestGetTransactionController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "failed get transaction by id",
			path:       "/transactions/1",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetTransactionControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

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

func TestGetTransactionsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get transactions",
			path:       "/transactions",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetTransactionsControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
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

func TestGetPatientTransactionsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "success get patient transactions",
			path:       "/transactions/patient/1",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.GetPatientTransactionsControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
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

func TestCreatePaymentController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "failed get patient transactions",
			path:       "/transactions/patient/1",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.CreatePaymentControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
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

func TestRescheduleController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "failed get reschedule",
			path:       "/transactions/1/reschedule",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.RescheduleControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestRescheduleController_invalid(t *testing.T) {
	e := echo.New()

	transactionID := uuid.New()

	rescheduleRequest := dto.ConsultationRescheduleRequest{
		Date: time.Now(),
		Session: "siang",
	}

	paymentJSON, err := json.Marshal(rescheduleRequest)
	assert.NoError(t, err, "Error marshalling reschedule request to JSON")

	req := httptest.NewRequest(http.MethodPost, "/transactions/"+transactionID.String(), bytes.NewBuffer(paymentJSON))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions" + transactionID.String() + "/reschedule")

	err = controller.RescheduleController(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response body")
	assert.Equal(t, "error parse id", response["message"])
}

func TestCancelTransactionController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "failed get cancel transaction",
			path:       "/transactions/1/cancel",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.CancelTransactionControllerTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestCancelTransactionController_invalid(t *testing.T) {
	e := echo.New()

	transactionID := uuid.New()

	rescheduleRequest := dto.TransactionRequest{
		Status:         "cancelled",
	}

	paymentJSON, err := json.Marshal(rescheduleRequest)
	assert.NoError(t, err, "Error marshalling reschedule request to JSON")

	req := httptest.NewRequest(http.MethodPost, "/transactions/"+transactionID.String(), bytes.NewBuffer(paymentJSON))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions" + transactionID.String() + "/cancel")

	err = controller.CancelTransactionController(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response body")
	assert.Equal(t, "error parse id", response["message"])
}

func TestValidateRefund(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "failed get validate refund",
			path:       "/refund/1",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.ValidateRefundTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}

func TestValidateRefund_invalid(t *testing.T) {
	e := echo.New()

	transactionID := uuid.New()

	validateRequest := dto.TransactionRequest{
		PaymentStatus: "refund",
	}

	paymentJSON, err := json.Marshal(validateRequest)
	assert.NoError(t, err, "Error marshalling reschedule request to JSON")

	req := httptest.NewRequest(http.MethodPost, "/transactions/"+transactionID.String(), bytes.NewBuffer(paymentJSON))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions" + transactionID.String() + "/cancel")

	err = controller.ValidateRefund(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Error unmarshalling response body")
	assert.Equal(t, "error parse id", response["message"])
}

func TestPaymentTimeOut(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "failed get payment timeout",
			path:       "/transactions/1/payment-timeout",
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataTransaction()
	token, user := InsertDataUser()

	for _, testCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)

		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath(testCase.path)
		context.SetParamNames("id")
		context.SetParamValues(user.ID.String())

		middleware.JWT([]byte(config.JWT_KEY))(controller.PaymentTimeOutTesting())(context)

		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		t.Run(fmt.Sprintf("GET %s", testCase.path), func(t *testing.T) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		})
	}
}