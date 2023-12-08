package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorGetSchedule struct {
	ID          uuid.UUID       `json:"id"`
	DoctorID    uuid.UUID       `json:"doctor_id"`
	PatientID   uuid.UUID       `json:"patient_id"`
	Date        time.Time       `json:"date"`
	Session     string          `json:"session"`
	QueueNumber string          `json:"queue_number"`
	Patient     PatientResponse `json:"patient"`
}

func ConvertToDoctorScheduleResponse(consultation model.Consultation, patient model.Patient) DoctorGetSchedule {
	return DoctorGetSchedule{
		ID:          consultation.ID,
		DoctorID:    consultation.DoctorID,
		PatientID:   consultation.PatientID,
		Date:        consultation.Date,
		QueueNumber: consultation.QueueNumber,
		Session:     consultation.Session,
		Patient:     ConvertToPatientResponse(patient),
	}
}
