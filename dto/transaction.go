package dto

import (
	"time"

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
	Payment        PaymentRequest    `json:"payment" form:"payment"`
}

type TransactionResponse struct {
	ID             uuid.UUID            `json:"id"`
	ConsultationID uuid.UUID            `json:"consultation_id"`
	Consultation   ConsultationResponse `json:"consultation"`
	Invoice        string               `json:"invoice"`
	Price          float64              `json:"price"`
	AdminPrice     float64              `json:"admin_price"`
	Total          float64              `json:"total"`
	Status         string               `json:"status"`
	PaymentStatus  string               `json:"payment_status"`
	CreatedAt	   time.Time 			`json:"created_at"`
	Refund         []RefundResponse       `json:"refund"`
	Payment        []PaymentResponse      `json:"payment"`
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
		PaymentStatus:  transaction.PaymentStatus,
	}
}

func ConvertToTransactionResponse(transaction model.Transaction) TransactionResponse {
	var articleRefundResponses []RefundResponse
	var articlePaymentResponses []PaymentResponse

	if transaction.Refund.ID != uuid.Nil {
		articleRefundResponses = append(articleRefundResponses, ConvertToRefundResponse(transaction.Refund))
	}
	if transaction.Payment.ID != uuid.Nil {
		articlePaymentResponses = append(articlePaymentResponses, ConvertToPaymentResponse(transaction.Payment))
	}

	return TransactionResponse{
		ID:             transaction.ID,
		ConsultationID: transaction.ConsultationID,
		Consultation:   ConvertToConsultationResponse(transaction.Consultation),
		Invoice:        transaction.Invoice,
		Price:          transaction.Price,
		AdminPrice:     transaction.AdminPrice,
		Total:          transaction.Total,
		Status:         string(transaction.Status),
		PaymentStatus:  string(transaction.PaymentStatus),
		Refund:        	articleRefundResponses,
		Payment:        articlePaymentResponses,
		CreatedAt: 		transaction.CreatedAt,
	}
}

func ConvertToTransactionDashboardResponse(transaction model.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:             transaction.ID,
		ConsultationID: transaction.ConsultationID,
	}
}
