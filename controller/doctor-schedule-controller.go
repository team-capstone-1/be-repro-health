package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"capstone-project/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllDoctorScheduleController(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	session := c.FormValue("session")
	dateString := c.FormValue("date")

	if dateString != "" {
		_, err := time.Parse("02-01-2006", dateString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed to parse date",
				"response": err.Error(),
			})
		}
	}

	responseData, err := repository.DoctorGetAllSchedules(doctor, session, dateString)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "failed get schedules data",
			"response": err.Error(),
		})
	}

	doctorSchedules := dto.ConvertToDoctorScheduleResponse(doctor, responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get schedule data",
		"response": doctorSchedules,
	})
}

func DoctorInactiveScheduleController(c echo.Context) error {
	doctorID := m.ExtractTokenUserId(c)
	if doctorID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	dateString := c.QueryParam("date")
	session := c.QueryParam("session")

	if session != "pagi" && session != "siang" && session != "malam" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid session value",
			"response": "session must be 'pagi', 'siang', or 'malam'.",
		})
	}

	doctorHoliday, err := repository.DoctorInactiveSchedule(doctorID, dateString, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed to mark doctor as inactive",
			"response": err.Error(),
		})
	}

	doctorHolidayResponse := HandleDoctorAction(doctorHoliday, err, doctorID, dateString, session)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success, doctor now not available",
		"response": doctorHolidayResponse,
	})
}

func DoctorActiveScheduleController(c echo.Context) error {
	doctorID := m.ExtractTokenUserId(c)
	if doctorID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	dateString := c.QueryParam("date")
	session := c.QueryParam("session")

	if session != "pagi" && session != "siang" && session != "malam" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "invalid session value",
			"response": "session must be 'pagi', 'siang', or 'malam'.",
		})
	}

	// Activate the doctor's schedule
	doctorHoliday, err := repository.DoctorActiveSchedule(doctorID, dateString, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed to mark doctor as active",
			"response": err.Error(),
		})
	}
	// Convert the updated schedule to the response format
	doctorHolidayResponse := HandleDoctorAction(doctorHoliday, err, doctorID, dateString, session)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success, doctor now available",
		"response": doctorHolidayResponse,
	})
}

func HandleDoctorAction(doctorHoliday model.Consultation, err error, doctorID uuid.UUID, date, session string) dto.DoctorScheduleResponse {
	if err != nil {
		return dto.DoctorScheduleResponse{
			DoctorID: doctorID,
			Data:     nil,
		}
	}

	// Fetch all schedules for the doctor
	responseData, err := repository.DoctorGetAllSchedules(doctorID, session, date)
	if err != nil {
		return dto.DoctorScheduleResponse{
			DoctorID: doctorID,
			Data:     nil,
		}
	}

	// Convert fetched schedules to the desired response format
	doctorSchedules := dto.ConvertToDoctorScheduleResponse(doctorID, responseData)

	// Find the date and session in the fetched schedules
	var foundDate bool
	var foundSession bool

	for i, data := range doctorSchedules.Data {
		if data.Date == date {
			foundDate = true

			for j, listData := range data.ListData {
				if listData.Session == session {
					foundSession = true

					// Update doctor availability based on DoctorHoliday status
					doctorSchedules.Data[i].ListData[j].DoctorAvailable = !doctorHoliday.DoctorAvailable

					// Set appointments to nil since doctor is not available
					doctorSchedules.Data[i].ListData[j].Appointments = nil
					break
				}
			}

			break
		}
	}

	// If the date or session is not found, add a new entry
	if !foundDate {
		// Create a new entry for the specified date
		newEntry := dto.FrontendData{
			Date: date,
			ListData: []dto.ListDetail{
				{
					DoctorAvailable: !doctorHoliday.DoctorAvailable,
					Session:         session,
					Appointments:    nil,
				},
			},
		}

		doctorSchedules.Data = append(doctorSchedules.Data, newEntry)
	} else if !foundSession {
		// Create a new entry for the specified session within the found date
		for i, data := range doctorSchedules.Data {
			if data.Date == date {
				newListData := dto.ListDetail{
					DoctorAvailable: !doctorHoliday.DoctorAvailable,
					Session:         session,
					Appointments:    nil,
				}

				doctorSchedules.Data[i].ListData = append(doctorSchedules.Data[i].ListData, newListData)
				break
			}
		}
	}

	return doctorSchedules
}

