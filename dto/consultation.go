package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type ConsultationRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id" form:"doctor_id"`
	PatientID uuid.UUID `json:"patient_id" form:"patient_id"`
	Date      time.Time `json:"date" form:"date"`
	Session   string    `json:"session" form:"session"`
}

type ConsultationRescheduleRequest struct {
	Date      time.Time `json:"date" form:"date"`
	Session   string    `json:"session" form:"session"`
}

type UserConsultationResponse struct {
	ID          uuid.UUID `json:"id"`
	DoctorID    uuid.UUID `json:"doctor_id"`
	PatientID   uuid.UUID `json:"patient_id"`
	ClinicID    uuid.UUID `json:"clinic_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Date        time.Time `json:"date"`
	Session     string    `json:"session"`
	QueueNumber string    `json:"queue_number"`
	Clinic      ClinicResponse    `json:"clinic"`
	Doctor      TransactionDoctorResponse    `json:"doctor"`
}

type ConsultationResponse struct {
	ID          uuid.UUID `json:"id"`
	DoctorID    uuid.UUID `json:"doctor_id"`
	PatientID   uuid.UUID `json:"patient_id"`
	ClinicID    uuid.UUID `json:"clinic_id"`
	Date        time.Time `json:"date"`
	Session     string    `json:"session"`
	QueueNumber string    `json:"queue_number"`
	Patient     PatientResponse   `json:"patient"`
	Clinic      ClinicResponse    `json:"clinic"`
	Doctor      TransactionDoctorResponse    `json:"doctor"`
}

func ConvertToConsultationModel(consultation ConsultationRequest) model.Consultation {
	return model.Consultation{
		ID:        uuid.New(),
		DoctorID:  consultation.DoctorID,
		PatientID: consultation.PatientID,
		Date:      consultation.Date,
		Session:   consultation.Session,
	}
}

func ConvertToConsultationRescheduleModel(consultation ConsultationRescheduleRequest, id uuid.UUID) model.Consultation {
	return model.Consultation{
		ID:        id,
		Date:      consultation.Date,
		Session:   consultation.Session,
	}
}

func ConvertToUserConsultationResponse(consultation model.Consultation) UserConsultationResponse {
	return UserConsultationResponse{
		ID:        consultation.ID,
		DoctorID:  consultation.DoctorID,
		PatientID: consultation.PatientID,
		ClinicID: consultation.ClinicID,
		Date: consultation.Date,
		QueueNumber: consultation.QueueNumber,
		Session: consultation.Session,
		Clinic: ConvertToClinicResponse(consultation.Clinic),
		Doctor: ConvertToTransactionDoctorResponse(consultation.Doctor),
	}
}

func ConvertToConsultationResponse(consultation model.Consultation) ConsultationResponse {
	return ConsultationResponse{
		ID:        consultation.ID,
		DoctorID:  consultation.DoctorID,
		PatientID: consultation.PatientID,
		ClinicID: consultation.ClinicID,
		Date: consultation.Date,
		QueueNumber: consultation.QueueNumber,
		Session: consultation.Session,
		Patient: ConvertToPatientResponse(consultation.Patient),
		Clinic: ConvertToClinicResponse(consultation.Clinic),
		Doctor: ConvertToTransactionDoctorResponse(consultation.Doctor),
	}
}