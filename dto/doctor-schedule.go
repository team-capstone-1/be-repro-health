package dto

import (
	"capstone-project/model"
	"log"

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
	Patient string    `json:"patient"`
	ID      uuid.UUID `json:"id"`
}

func ConvertToAppointments(patientResponses []PatientResponse) []Appointment {
	var appointments []Appointment

	for _, patientResponse := range patientResponses {
		// Assuming ConvertToPatientResponse returns a valid PatientResponse
		appointment := Appointment{
			Patient: patientResponse.Name, // Assuming Name is the patient's name
			ID:      patientResponse.ID,   // Assuming ID is the patient's ID
		}

		appointments = append(appointments, appointment)
	}

	return appointments
}

func ConvertToDoctorScheduleResponse(doctorID uuid.UUID, schedules []model.Consultation) DoctorScheduleResponse {
	// Group schedules by date and session
	doctorSchedulesMap := make(map[string]map[string][]PatientResponse)

	for _, schedule := range schedules {
		// Use the correct date format when constructing the key
		date := schedule.Date.Format("02-01-2006")
		session := schedule.Session

		// Retrieve or create the map for the current date
		patientResponsesMap, exists := doctorSchedulesMap[date]
		if !exists {
			patientResponsesMap = make(map[string][]PatientResponse)
			doctorSchedulesMap[date] = patientResponsesMap
		}

		// Retrieve or create the slice for the current session
		patientResponses := patientResponsesMap[session]
		if patientResponses == nil {
			patientResponses = make([]PatientResponse, 0)
		}

		// Access appointments directly from the Consultation object
		for _, appointment := range schedule.Transaction {
			// Add a log statement to check the content of appointment.Consultation.Patient
			log.Println("Patient data:", appointment.Consultation.Patient)

			// Add a log statement to check the content of ConvertToPatientResponse result
			patientResponse := ConvertToPatientResponse(appointment.Consultation.Patient)
			log.Println("PatientResponse:", patientResponse)

			// Access appointments directly from the Consultation object
			for _, appointment := range schedule.Transaction {
				// Assuming ConvertToPatientResponse returns a valid PatientResponse
				patientResponses = append(patientResponses, ConvertToPatientResponse(appointment.Consultation.Patient))
			}
		}

		// Update the map with the appointments for the current session
		patientResponsesMap[session] = patientResponses
	}

	// Convert the map to DoctorScheduleResponse format
	var doctorSchedules []FrontendData
	for date, patientResponsesMap := range doctorSchedulesMap {
		var listData []ListDetail

		// Separate morning, afternoon, and evening sessions
		for session, patientResponses := range patientResponsesMap {
			// Adjust the doctorAvailable based on your logic
			doctorAvailable := len(patientResponses) > 0

			// ConvertToAppointments converts a slice of PatientResponse to a slice of Appointment
			appointments := ConvertToAppointments(patientResponses)

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

// type DoctorScheduleResponse struct {
// 	DoctorID uuid.UUID        `json:"doctor_id"`
// 	Data     []DoctorSchedule `json:"data"`
// }

// type DoctorSchedule struct {
// 	Date     string        `json:"date"`
// 	ListData []Appointment `json:"listData"`
// }

// type Appointment struct {
// 	DoctorAvailable bool              `json:"doctor_available"`
// 	Session         string            `json:"session"`
// 	Appointments    []PatientResponse `json:"appointments"`
// }

// func ConvertToDoctorScheduleResponse(doctorID uuid.UUID, date string, schedules []model.Consultation) DoctorScheduleResponse {
// 	var doctorSchedulesMap = make(map[string][]PatientResponse)

// 	// Group schedules by date
// 	for _, schedule := range schedules {
// 		date := schedule.Date.Format("02-01-2006")
// 		patientResponses := doctorSchedulesMap[date]

// 		// Access appointments directly from the Consultation object
// 		for _, appointment := range schedule.Transaction {
// 			patientResponses = append(patientResponses, ConvertToPatientResponse(appointment.Consultation.Patient))
// 		}

// 		doctorSchedulesMap[date] = patientResponses
// 	}

// 	// Convert the map to DoctorScheduleResponse format
// 	var doctorSchedules []DoctorSchedule
// 	for date, patientResponses := range doctorSchedulesMap {
// 		doctorSchedules = append(doctorSchedules, DoctorSchedule{
// 			Date: date,
// 			ListData: []Appointment{
// 				{
// 					DoctorAvailable: len(patientResponses) > 0,
// 					Session:         schedules[0].Session, // Assuming the session is the same for all appointments on the same date
// 					Appointments:    patientResponses,
// 				},
// 			},
// 		})
// 	}

// 	return DoctorScheduleResponse{
// 		DoctorID: doctorID,
// 		Data:     doctorSchedules,
// 	}
// }
