package dto

import (
	"capstone-project/model"
	"fmt"

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
	ID      uuid.UUID `json:"patient_id"`
}

func ConvertToDoctorScheduleResponse(doctorID uuid.UUID, schedules []model.Consultation) DoctorScheduleResponse {
	// Group schedules by date and session
	doctorSchedulesMap := make(map[string]map[string][]model.Consultation)

	for _, schedule := range schedules {
		date := schedule.Date.Format("02-01-2006")
		session := schedule.Session

		// Retrieve or create the map for the current date
		consultationMap, exists := doctorSchedulesMap[date]
		if !exists {
			consultationMap = make(map[string][]model.Consultation)
			doctorSchedulesMap[date] = consultationMap
		}

		// Retrieve or create the slice for the current session
		consultations := consultationMap[session]
		if consultations == nil {
			consultations = make([]model.Consultation, 0)
		}

		// Append the current schedule to the slice
		consultations = append(consultations, schedule)

		// Update the map with the schedules for the current session
		consultationMap[session] = consultations
	}

	// Convert the map to DoctorScheduleResponse format
	var doctorSchedules []FrontendData
	for date, consultationMap := range doctorSchedulesMap {
		var listData []ListDetail

		// Separate morning, afternoon, and evening sessions
		for _, session := range []string{"pagi", "siang", "malam"} {
			// Check if there are any schedules for the current session
			consultations := consultationMap[session]

			// Assume the default value is true, update if there's a valid value in the schedule
			doctorAvailable := true

			// Iterate through the schedules to find the doctor availability
			for _, consultation := range consultations {
				patientResponse := ConvertToPatientResponse(consultation.Patient)
				// Assuming Consultation.ID is the appointment's ID
				appointment := Appointment{
					Patient: patientResponse.Name,
					ID:      consultation.PatientID,
				}
				fmt.Print(appointment)
				// Check if DoctorAvailable is false
				if !consultation.DoctorAvailable {
					doctorAvailable = false
					break
				}
			}

			// ConvertToAppointments converts a slice of Consultation to a slice of Appointment
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
		// Assuming ConvertToPatientResponse returns a valid PatientResponse
		patientResponse := ConvertToPatientResponse(consultation.Patient)
		// Assuming Consultation.ID is the appointment's ID
		appointment := Appointment{
			Patient: patientResponse.Name,
			ID:      patientResponse.ID,
		}

		appointments = append(appointments, appointment)
	}

	return appointments
}

// func ConvertToAppointments(patientResponses []PatientResponse) []Appointment {
// 	if patientResponses == nil {
// 		return nil
// 	}

// 	var appointments []Appointment

// 	for _, patientResponse := range patientResponses {
// 		// Assuming ConvertToPatientResponse returns a valid PatientResponse
// 		appointment := Appointment{
// 			Patient: patientResponse.Name, // Assuming Name is the patient's name
// 			ID:      patientResponse.ID,   // Assuming ID is the patient's ID
// 		}

// 		appointments = append(appointments, appointment)
// 	}

// 	return appointments
// }

// func ConvertToDoctorScheduleResponse(doctorID uuid.UUID, schedules []model.Consultation) DoctorScheduleResponse {
// 	// Group schedules by date and session
// 	doctorSchedulesMap := make(map[string]map[string][]PatientResponse)

// 	for _, schedule := range schedules {
// 		date := schedule.Date.Format("02-01-2006")
// 		session := schedule.Session

// 		// Retrieve or create the map for the current date
// 		patientResponsesMap, exists := doctorSchedulesMap[date]
// 		if !exists {
// 			patientResponsesMap = make(map[string][]PatientResponse)
// 			doctorSchedulesMap[date] = patientResponsesMap
// 		}

// 		// Retrieve or create the slice for the current session
// 		patientResponses := patientResponsesMap[session]
// 		if patientResponses == nil {
// 			patientResponses = make([]PatientResponse, 0)
// 		}

// 		// Access appointments directly from the Consultation object
// 		for _, appointment := range schedule.Transaction {
// 			patientResponse := ConvertToPatientResponse(appointment.Consultation.Patient)
// 			patientResponses = append(patientResponses, patientResponse)
// 		}

// 		// Update the map with the appointments for the current session
// 		patientResponsesMap[session] = patientResponses
// 	}

// 	// Convert the map to DoctorScheduleResponse format
// 	var doctorSchedules []FrontendData
// 	for date, patientResponsesMap := range doctorSchedulesMap {
// 		var listData []ListDetail

// 		// Separate morning, afternoon, and evening sessions
// 		for _, session := range []string{"pagi", "siang", "malam"} {
// 			// Check if there are any appointments for the current session
// 			patientResponses, exists := patientResponsesMap[session]
// 			doctorAvailable := exists

// 			if exists {
// 				// Check if there are any appointments for the current session
// 				doctorAvailable = len(patientResponses) > 0
// 			}

// 			// ConvertToAppointments converts a slice of PatientResponse to a slice of Appointment
// 			appointments := ConvertToAppointments(patientResponses)

// 			listData = append(listData, ListDetail{
// 				DoctorAvailable: doctorAvailable,
// 				Session:         session,
// 				Appointments:    appointments,
// 			})
// 		}

// 		doctorSchedules = append(doctorSchedules, FrontendData{
// 			Date:     date,
// 			ListData: listData,
// 		})
// 	}

// 	return DoctorScheduleResponse{
// 		DoctorID: doctorID,
// 		Data:     doctorSchedules,
// 	}
// }
