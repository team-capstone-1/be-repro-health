package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type ConsultationRequest struct {
	DoctorID      uuid.UUID `json:"doctor_id" form:"doctor_id"`
	PatientID     uuid.UUID `json:"patient_id" form:"patient_id"`
	Date          string    `json:"date" form:"date"`
	Session       string    `json:"session" form:"session"`
	PaymentMethod string    `json:"payment_method" form:"payment_method"`
}

type ConsultationRescheduleRequest struct {
	Date    time.Time `json:"date" form:"date"`
	Session string    `json:"session" form:"session"`
}

type UserConsultationResponse struct {
	ID            uuid.UUID                 `json:"id"`
	DoctorID      uuid.UUID                 `json:"doctor_id"`
	PatientID     uuid.UUID                 `json:"patient_id"`
	ClinicID      uuid.UUID                 `json:"clinic_id"`
	TransactionID uuid.UUID                 `json:"transaction_id"`
	Date          time.Time                 `json:"date"`
	Session       string                    `json:"session"`
	QueueNumber   string                    `json:"queue_number"`
	PaymentMethod string                    `json:"payment_method"`
	Rescheduled   bool                      `json:"rescheduled"`
	DoctorAvailable bool					`json:"doctor_available"`
	Clinic        ClinicResponse            `json:"clinic"`
	Doctor        TransactionDoctorResponse `json:"doctor"`
}

type ConsultationResponse struct {
	ID            uuid.UUID                 `json:"id"`
	DoctorID      uuid.UUID                 `json:"doctor_id"`
	PatientID     uuid.UUID                 `json:"patient_id"`
	ClinicID      uuid.UUID                 `json:"clinic_id"`
	Date          time.Time                 `json:"date"`
	Session       string                    `json:"session"`
	QueueNumber   string                    `json:"queue_number"`
	PaymentMethod string                    `json:"payment_method"`
	Rescheduled   bool                      `json:"rescheduled"`
	DoctorAvailable bool					`json:"doctor_available"`
	Patient       PatientResponse           `json:"patient"`
	Clinic        ClinicResponse            `json:"clinic"`
	Doctor        TransactionDoctorResponse `json:"doctor"`
}

func ConvertToConsultationModel(consultation ConsultationRequest) model.Consultation {
	parsedDate, _ := time.Parse("2006-01-02", consultation.Date)

	return model.Consultation{
		ID:            uuid.New(),
		DoctorID:      consultation.DoctorID,
		PatientID:     consultation.PatientID,
		Date:          parsedDate,
		Session:       consultation.Session,
		PaymentMethod: consultation.PaymentMethod,
	}
}

func ConvertToConsultationRescheduleModel(consultation ConsultationRescheduleRequest, id uuid.UUID) model.Consultation {
	return model.Consultation{
		ID:          id,
		Date:        consultation.Date,
		Session:     consultation.Session,
		Rescheduled: true,
	}
}

func ConvertToUserConsultationResponse(consultation model.Consultation) UserConsultationResponse {
	return UserConsultationResponse{
		ID:            consultation.ID,
		DoctorID:      consultation.DoctorID,
		PatientID:     consultation.PatientID,
		ClinicID:      consultation.ClinicID,
		Date:          consultation.Date,
		QueueNumber:   consultation.QueueNumber,
		PaymentMethod: consultation.PaymentMethod,
		Session:       consultation.Session,
		Rescheduled:   consultation.Rescheduled,
		DoctorAvailable: consultation.DoctorAvailable,
		Clinic:        ConvertToClinicResponse(consultation.Clinic),
		Doctor:        ConvertToTransactionDoctorResponse(consultation.Doctor),
	}
}

func ConvertToConsultationResponse(consultation model.Consultation) ConsultationResponse {
	return ConsultationResponse{
		ID:            consultation.ID,
		DoctorID:      consultation.DoctorID,
		PatientID:     consultation.PatientID,
		ClinicID:      consultation.ClinicID,
		Date:          consultation.Date,
		QueueNumber:   consultation.QueueNumber,
		PaymentMethod: consultation.PaymentMethod,
		Session:       consultation.Session,
		Rescheduled:   consultation.Rescheduled,
		DoctorAvailable: consultation.DoctorAvailable,
		Patient:       ConvertToPatientResponse(consultation.Patient),
		Clinic:        ConvertToClinicResponse(consultation.Clinic),
		Doctor:        ConvertToTransactionDoctorResponse(consultation.Doctor),
	}
}
