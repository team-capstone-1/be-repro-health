package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorPatientConsulationResponse struct {
	ID        uuid.UUID       `json:"id"`
	PatientID uuid.UUID       `json:"patient_id"`
	Patient   PatientResponse `json:"patient"`
	Date      time.Time       `json:"date"`
	Session   string          `json:"session"`
}

type DoctorTransactionResponse struct {
	ID      uuid.UUID           `json:"id"`
	Payment TransactionResponse `json:"payment"`
	Price   float64             `json:"price"`
	Status  string              `json:"status"`
}

type DoctorPaymentResponse struct {
	ID     uuid.UUID       `json:"id"`
	Method PaymentResponse `json:"payment_method"`
}

type DoctorGetDetailsTransactionResponse struct {
	ID            uuid.UUID `json:"id"`
	Invoice       string    `json:"invoice"`
	Date          time.Time `json:"date"`
	PaymentMethod string    `json:"payment_method"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	AdminFee      float64   `json:"admin_fee"`
	Total         float64   `json:"total"`
}

func ConvertToDoctorGetDetailsTransactionResponse(transaction model.Transaction) DoctorGetDetailsTransactionResponse {
	return DoctorGetDetailsTransactionResponse{
		ID:            transaction.ID,
		Invoice:       transaction.Invoice,
		Date:          transaction.Date,
		PaymentMethod: transaction.Payment.Method,
		Name:          transaction.Payment.Name,
		Price:         transaction.Price,
		AdminFee:      transaction.AdminPrice,
		Total:         transaction.Total,
	}
}

func ConvertToDoctorPatientConsultationResponse(consultation model.Consultation) DoctorPatientConsulationResponse {

	return DoctorPatientConsulationResponse{
		ID:        consultation.ID,
		PatientID: consultation.PatientID,
		Date:      consultation.Date,
		Session:   consultation.Session,
	}
}

func ConvertToDoctorTransactionResponse(transaction model.Transaction) DoctorTransactionResponse {
	return DoctorTransactionResponse{
		ID:      transaction.ID,
		Payment: ConvertToTransactionResponse(transaction),
		Price:   transaction.Price,
		Status:  string(transaction.Status),
	}
}

func ConvertToDoctorPaymentResponse(payment model.Payment) DoctorPaymentResponse {
	return DoctorPaymentResponse{
		ID:     payment.ID,
		Method: ConvertToPaymentResponse(payment),
	}
}
