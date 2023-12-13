package dto

import (
	"capstone-project/database"
	"capstone-project/model"
	"capstone-project/repository"
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
	// Create a map to store consultations based on date and session
	doctorSchedulesMap := make(map[string]map[string][]model.Consultation)

	for _, schedule := range schedules {
		date := schedule.Date.Format("2006-01-02")
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

	if len(doctorSchedulesMap) > 0 {
		for date, consultationMap := range doctorSchedulesMap {
			var listData []ListDetail

			for _, session := range []string{"pagi", "siang", "malam"} {
				consultations := consultationMap[session]

				doctorAvailable := true

				var appointments []Appointment // Declare appointments variable here

				if len(consultations) > 0 {
					for _, consultation := range consultations {
						patientResponse := ConvertToPatientResponse(consultation.Patient)
						appointment := Appointment{
							ConsultationID: consultation.ID,
							Patient:        patientResponse.Name,
							PatientID:      consultation.PatientID,
						}

						appointments = append(appointments, appointment)

						if !consultation.DoctorAvailable {
							doctorAvailable = false
						}
					}
				}

				// Check if Doctor is on holiday for the given date and session
				isDoctorOnHoliday, err := repository.IsDoctorOnHoliday(doctorID, date, session)
				if err != nil {
					// Handle the error if needed
					fmt.Println("Error checking doctor's holiday:", err)
				}

				// Update doctorAvailable based on DoctorHoliday table
				if isDoctorOnHoliday {
					doctorAvailable = false
				}

				// Include a default entry even if there are no appointments
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
	}

	return DoctorScheduleResponse{
		DoctorID: doctorID,
		Data:     doctorSchedules,
	}
}

func GetDoctorHolidayResponse(doctorID uuid.UUID, date string, sessions []string) (DoctorScheduleResponse, error) {
	var doctorHolidays []model.DoctorHoliday

	// Fetch DoctorHoliday data for the given date and sessions
	err := database.DB.Where("doctor_id = ? AND date = ? AND session IN ?", doctorID, date, sessions).
		Find(&doctorHolidays).Error

	if err != nil {
		return DoctorScheduleResponse{}, err
	}

	// Create a map to store doctor holidays based on date and session
	doctorHolidaysMap := make(map[string]map[string][]model.DoctorHoliday)

	for _, holiday := range doctorHolidays {
		holidayDate := holiday.Date.Format("02-01-2006")
		holidaySession := holiday.Session

		holidayMap, exists := doctorHolidaysMap[holidayDate]
		if !exists {
			holidayMap = make(map[string][]model.DoctorHoliday)
			doctorHolidaysMap[holidayDate] = holidayMap
		}

		holidays := holidayMap[holidaySession]
		if holidays == nil {
			holidays = make([]model.DoctorHoliday, 0)
		}

		holidays = append(holidays, holiday)

		holidayMap[holidaySession] = holidays
	}

	// Format the response similar to DoctorScheduleResponse
	var doctorSchedules []FrontendData

	if len(doctorHolidaysMap) > 0 {
		for holidayDate := range doctorHolidaysMap {
			var listData []ListDetail

			for _, session := range sessions {
				// holidays := holidayMap[session]

				doctorAvailable := false // Set to false since it's a doctor holiday

				var appointments []Appointment // Appointments can be empty for holidays

				// Include a default entry even if there are no appointments
				listData = append(listData, ListDetail{
					DoctorAvailable: doctorAvailable,
					Session:         session,
					Appointments:    appointments,
				})
			}

			doctorSchedules = append(doctorSchedules, FrontendData{
				Date:     holidayDate,
				ListData: listData,
			})
		}
	}

	return DoctorScheduleResponse{
		DoctorID: doctorID,
		Data:     doctorSchedules,
	}, nil
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