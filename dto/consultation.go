package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type ConsultationRequest struct {
	PatientID uuid.UUID `json:"patient_id" form:"patient_id"`
	Date      time.Time `json:"date" form:"date"`
	Session   string    `json:"session" form:"session"`
}

type ConsultationResponse struct {
	ID          uuid.UUID `json:"id"`
	DoctorID    uuid.UUID `json:"doctor_id"`
	PatientID   uuid.UUID `json:"patient_id"`
	ClinicID    uuid.UUID `json:"clinic_id"`
	Date        time.Time `json:"date"`
	Session     string    `json:"session"`
	Clinic      ClinicResponse    `json:"clinic"`
	Doctor      DoctorResponse    `json:"doctor"`
}

func ConvertToConsultationModel(consultation ConsultationRequest) model.Consultation {
	return model.Consultation{
		ID:        uuid.New(),
		PatientID: consultation.PatientID,
		Date:      consultation.Date,
		Session:   consultation.Session,
	}
}

func ConvertToConsultationResponse(consultation model.Consultation) ConsultationResponse {
	return ConsultationResponse{
		ID:        consultation.ID,
		DoctorID:  consultation.DoctorID,
		PatientID: consultation.PatientID,
		ClinicID: consultation.ClinicID,
		Date: consultation.Date,
		Session: consultation.Session,
		Clinic: ConvertToClinicResponse(consultation.Clinic),
		Doctor: ConvertToDoctorResponse(consultation.Doctor),
	}
}
