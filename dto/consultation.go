package dto

import (
	"time"
	"github.com/google/uuid"
	"capstone-project/model"
)

type ConsultationRequest struct {
	// DoctorID    uuid.UUID `json:"doctor_id" form:"doctor_id"`
	// PatientID   uuid.UUID `json:"consultation_id" form:"consultation_id"`
	ClinicID    uuid.UUID `json:"clinic_id" form:"clinic_id"`
	Date        time.Time `json:"date" form:"date"`
	Session     string    `json:"session" form:"session"`
}

type ConsultationResponse struct {
	ID          uuid.UUID `json:"id"`
	// DoctorID    uuid.UUID `json:"doctor_id"`
	// PatientID   uuid.UUID `json:"consultation_id"`
	ClinicID    uuid.UUID `json:"clinic_id"`
	Date        time.Time `json:"date"`
	Session     string    `json:"session"`
}

func ConvertToConsultationModel(consultation ConsultationRequest) model.Consultation {
	return model.Consultation{
		ID: uuid.New(),
		// DoctorID: consultation.DoctorID,
		// PatientID: consultation.PatientID,
		ClinicID: consultation.ClinicID,
		Date: consultation.Date,
		Session: consultation.Session,
	}
}

func ConvertToConsultationResponse(consultation model.Consultation) ConsultationResponse {
	return ConsultationResponse{
		ID:    consultation.ID,
		// DoctorID: consultation.DoctorID,
		// PatientID: consultation.PatientID,
		ClinicID: consultation.ClinicID,
		Date: consultation.Date,
		Session: consultation.Session,
	}
}