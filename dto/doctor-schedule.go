package dto

import (
	"capstone-project/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DoctorScheduleResponse struct {
	DoctorID uuid.UUID      `json:"doctor_id"`
	Data     []FrontendData `json:"data"`
}

type FrontendData struct {
	Date     string       `json:"date"`
	ListData []ListDetail `json:"listData"`
}

type ListDetail struct {
	DoctorAvailable bool          `json:"doctor_available"`
	Session         string        `json:"session"`
	Appointments    []Appointment `json:"appointments"`
}

type Appointment struct {
	ConsultationID uuid.UUID `json:"consultation_id"`
	Patient        string    `json:"patient"`
	PatientID      uuid.UUID `json:"patient_id"`
}

func ConvertToDoctorScheduleResponse(doctorID uuid.UUID, schedules []model.Consultation) DoctorScheduleResponse {
	doctorSchedulesMap := make(map[string]map[string][]model.Consultation)

	for _, schedule := range schedules {
		date := schedule.Date.Format("02-01-2006")
		session := schedule.Session

		consultationMap, exists := doctorSchedulesMap[date]
		if !exists {
			consultationMap = make(map[string][]model.Consultation)
			doctorSchedulesMap[date] = consultationMap
		}

		consultations := consultationMap[session]
		if consultations == nil {
			consultations = make([]model.Consultation, 0)
		}

		consultations = append(consultations, schedule)

		consultationMap[session] = consultations
	}

	var doctorSchedules []FrontendData
	for date, consultationMap := range doctorSchedulesMap {
		var listData []ListDetail

		for _, session := range []string{"pagi", "siang", "malam"} {
			consultations := consultationMap[session]

			doctorAvailable := true

			for _, consultation := range consultations {
				patientResponse := ConvertToPatientResponse(consultation.Patient)
				appointment := Appointment{
					ConsultationID: consultation.ID,
					Patient:        patientResponse.Name,
					PatientID:      consultation.PatientID,
				}
				fmt.Print(appointment)

				if !consultation.DoctorAvailable {
					doctorAvailable = false
					break
				}
			}

			appointments := ConvertToAppointments(consultations)

			listData = append(listData, ListDetail{
				DoctorAvailable: doctorAvailable,
				Session:         session,
				Appointments:    appointments,
			})
		}

		doctorSchedules = append(doctorSchedules, FrontendData{
			Date:     date,
			ListData: listData,
		})
	}

	return DoctorScheduleResponse{
		DoctorID: doctorID,
		Data:     doctorSchedules,
	}
}

func ConvertToAppointments(consultations []model.Consultation) []Appointment {
	var appointments []Appointment

	for _, consultation := range consultations {
		patientResponse := ConvertToPatientResponse(consultation.Patient)
		appointment := Appointment{
			ConsultationID: consultation.ID,
			Patient:        patientResponse.Name,
			PatientID:      patientResponse.ID,
		}

		appointments = append(appointments, appointment)
	}

	return appointments
}

// DOCTOR HOLIDAY

type DoctorHolidayRequest struct {
	DoctorAvailable bool `json:"doctor_available"`
}

func ConvertToModelDoctorHoliday(doctorHoliday DoctorHolidayRequest) model.Consultation {
	return model.Consultation{
		DoctorAvailable: doctorHoliday.DoctorAvailable,
	}
}

type DoctorHolidayResponse struct {
	ID              uuid.UUID `json:"id"`
	DoctorID        uuid.UUID `json:"doctor_id"`
	Date            time.Time `json:"date"`
	Session         string    `json:"session"`
	DoctorAvailable bool      `json:"doctor_available"`
}

func ConvertToDoctorHolidayResponse(doctorHoliday model.Consultation) DoctorHolidayResponse {
	return DoctorHolidayResponse{
		ID:              doctorHoliday.ID,
		DoctorID:        doctorHoliday.DoctorID,
		Date:            doctorHoliday.Date,
		Session:         doctorHoliday.Session,
		DoctorAvailable: doctorHoliday.DoctorAvailable,
	}
}
