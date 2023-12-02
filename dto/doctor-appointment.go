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
