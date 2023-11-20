package dto

import (
	"capstone-project/model"

	"github.com/google/uuid"
)

type TransactionRequest struct {
	ConsultationID uuid.UUID `json:"consultation_id" form:"consultation_id"`
	Invoice        string    `json:"invoice" form:"invoice"`
	Price          float64   `json:"price" form:"price"`
	AdminPrice     float64   `json:"admin_price" form:"admin_price"`
	Total          float64   `json:"total" form:"total"`
	Status         string    `json:"status" form:"status"`
	PaymentStatus  string    `json:"payment_status" form:"payment_status"`
	Refunds        float64   `json:"refunds" form:"refunds"`
	PaymentMethods string    `json:"payment_methods" form:"payment_methods"`
}

type TransactionResponse struct {
	ID             uuid.UUID `json:"id"`
	ConsultationID uuid.UUID `json:"consultation_id"`
	Invoice        string    `json:"invoice"`
	Price          float64   `json:"price"`
	AdminPrice     float64   `json:"admin_price"`
	Total          float64   `json:"total"`
	Status         string    `json:"status"`
	PaymentStatus  string    `json:"payment_status"`
	Refunds        float64   `json:"refunds"`
	PaymentMethods string    `json:"payment_methods"`
}

func ConvertToTransactionModel(transaction TransactionRequest) model.Transaction {
	return model.Transaction{
		ID:             uuid.New(),
		ConsultationID: transaction.ConsultationID,
		Invoice:        transaction.Invoice,
		Price:          transaction.Price,
		AdminPrice:     transaction.AdminPrice,
		Total:          transaction.Total,
		Status:         model.TransactionStatus(transaction.Status),
		PaymentStatus:  model.TransactionStatus(transaction.PaymentStatus),
	}
}

func ConvertToTransactionDashboardResponse(transaction model.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:             transaction.ID,
		ConsultationID: transaction.ConsultationID,
	}
}
