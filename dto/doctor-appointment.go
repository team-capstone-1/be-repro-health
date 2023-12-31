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
	TransactionID   uuid.UUID `json:"transaction_id"`
	PatientName     string    `json:"patient_name"`
	ProfileImage    string    `json:"profile_image"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Gender          string    `json:"gender"`
	Height          float64   `json:"height"`
	Weight          float64   `json:"weight"`
	TelephoneNumber string    `json:"telephone_number"`
	Email           string    `json:"email"`
	QueueNumber     string    `json:"sequence_number"`
	Date            time.Time `json:"date"`
	Session         string    `json:"session"`
	Location        string    `json:"location"`
	PaymentMethod   string    `json:"payment_method"`
	Total           float64   `json:"total"`
	Status          string    `json:"status"`
}

type DoctorGetAllConsultations struct {
	ID             uuid.UUID `json:"id"`
	PatientID      uuid.UUID `json:"patient_id"`
	PatientName    string    `json:"patient_name"`
	SequenceNumber string    `json:"sequence_number"`
	Date           time.Time `json:"date"`
	Session        string    `json:"session"`
	Total          float64   `json:"total"`
	PaymentMethod  string    `json:"payment_method"`
	Status         string    `json:"status"`
}

type DoctorConfirmConsultationRequest struct {
	ConsultationID uuid.UUID `json:"consultation_id"`
}

func ConvertToDoctorConfirmConsultationModel(consultation DoctorConfirmConsultationRequest, consultationID uuid.UUID) model.Transaction {
	return model.Transaction{
		ConsultationID: consultationID,
	}
}
func ConvertToDoctorFinishConsultationModel(consultation DoctorConfirmConsultationRequest, consultationID uuid.UUID) model.Transaction {
	return model.Transaction{
		ConsultationID: consultationID,
	}
}

func ConvertToDoctorGetAllConsultations(consultation model.Consultation) DoctorGetAllConsultations {

	var total float64
	var paymentMethod string
	var status string

	for i := range consultation.Transaction {

		total = consultation.Transaction[i].Total
		paymentMethod = consultation.PaymentMethod
		status = string(consultation.Transaction[i].Status)
	}

	return DoctorGetAllConsultations{
		ID:             consultation.ID,
		PatientID:      consultation.PatientID,
		PatientName:    consultation.Patient.Name,
		SequenceNumber: consultation.QueueNumber,
		Date:           consultation.Date,
		Session:        consultation.Session,
		Total:          total,
		PaymentMethod:  paymentMethod,
		Status:         status,
	}
}

func ConvertToDoctorGetDetailsPatientResponse(consultation model.Consultation) DoctorGetDetailsPatientResponse {
	var transactionID uuid.UUID
	var paymentMethod string
	var total float64
	var status string

	if len(consultation.Transaction) > 0 {
		transactionID = consultation.Transaction[0].ID
		paymentMethod = consultation.PaymentMethod
		total = consultation.Transaction[0].Total
		status = string(consultation.Transaction[0].Status)
	}

	return DoctorGetDetailsPatientResponse{
		ID:              consultation.ID,
		PatientID:       consultation.PatientID,
		TransactionID:   transactionID,
		PatientName:     consultation.Patient.Name,
		ProfileImage:    consultation.Patient.ProfileImage,
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
		QueueNumber:     consultation.QueueNumber,
		Total:           total,
		Status:          status,
	}
}

func ConvertToDoctorGetDetailsTransactionResponse(transaction model.Transaction) DoctorGetDetailsTransactionResponse {
	return DoctorGetDetailsTransactionResponse{
		ID:            transaction.ID,
		Invoice:       transaction.Invoice,
		Date:          transaction.CreatedAt,
		PaymentMethod: transaction.Consultation.PaymentMethod,
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
