package dto

import (
	"capstone-project/model"
	"strings"
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
	Invoice string              `json:"invoice"`
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

type DoctorGetDetailsPatientResponse struct {
	ID              uuid.UUID `json:"id"`
	PatientID       uuid.UUID `json:"patient_id"`
	PatientName     string    `json:"patient_name"`
	ProfileImage    string    `json:"profile_image"`
	TransactionID   uuid.UUID `json:"transaction_id"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Gender          string    `json:"gender"`
	Height          float64   `json:"height"`
	Weight          float64   `json:"weight"`
	TelephoneNumber string    `json:"telephone_number"`
	Email           string    `json:"email"`
	SequenceNumber  string    `json:"sequence_number"`
	Date            time.Time `json:"date"`
	Session         string    `json:"session"`
	Location        string    `json:"location"`
	PaymentMethod   string    `json:"payment_method"`
	Total           float64   `json:"total"`
	Status          string    `json:"status"`
}

func ConvertToDoctorGetDetailsPatientResponse(consultation model.Consultation) DoctorGetDetailsPatientResponse {
	var transactionID uuid.UUID
	var paymentMethod string
	var total float64
	var invoice string
	var status string

	if len(consultation.Transaction) > 0 {
		transactionID = consultation.Transaction[0].ID
		paymentMethod = consultation.Transaction[0].Payment.Method
		total = consultation.Transaction[0].Total
		invoice = consultation.Transaction[0].Invoice
		status = string(consultation.Transaction[0].Status)
	}

	parts := strings.Split(invoice, "/")
	sequenceNumber := parts[len(parts)-1]

	return DoctorGetDetailsPatientResponse{
		ID:              consultation.ID,
		PatientID:       consultation.PatientID,
		PatientName:     consultation.Patient.Name,
		ProfileImage:    consultation.Patient.ProfileImage,
		TransactionID:   transactionID,
		DateOfBirth:     consultation.Patient.DateOfBirth,
		Gender:          consultation.Patient.Gender,
		Height:          consultation.Patient.Height,
		Weight:          consultation.Patient.Weight,
		TelephoneNumber: consultation.Patient.TelephoneNumber,
		Email:           consultation.Patient.User.Email,
		Date:            consultation.Date,
		Session:         consultation.Session,
		Location:        consultation.Clinic.Location,
		PaymentMethod:   paymentMethod,
		SequenceNumber:  sequenceNumber,
		Total:           total,
		Status:          status,
	}
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
		Invoice: transaction.Invoice,
		Status:  string(transaction.Status),
	}
}

func ConvertToDoctorPaymentResponse(payment model.Payment) DoctorPaymentResponse {
	return DoctorPaymentResponse{
		ID:     payment.ID,
		Method: ConvertToPaymentResponse(payment),
	}
}
