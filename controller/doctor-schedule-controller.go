package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
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
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	session := c.QueryParam("session")
	dateString := c.FormValue("date")

	var date time.Time
	if dateString != "" {
		parsedDate, err := time.Parse("02-01-2006", dateString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message":  "failed to parse date",
				"response": err.Error(),
			})
		}
		date = parsedDate
	}

	if session != "pagi" && session != "siang" && session != "malam" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "invalid session value",
			"response": "Session must be 'pagi', 'siang', or 'malam'.",
		})
	}

	doctorHoliday := dto.DoctorHolidayRequest{
		ID:       uuid.New(),
		DoctorID: doctor,
		Date:     date,
	}
	if err := c.Bind(&doctorHoliday); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "failed to bind data",
			"response": err.Error(),
		})
	}

	modelDoctorHoliday := dto.ConvertToModelDoctorHoliday(doctorHoliday)

	// Check if there are existing doctor holidays for the specified doctor, date, and session
	existingHolidays, err := repository.GetDoctorHolidaysByDateAndSession(doctor, date, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":  "failed to check existing doctor holidays",
			"response": err.Error(),
		})
	}

	// Get consultations for the specified doctor, date, and session
	consultations, err := repository.GetConsultationsByDoctorSchedule(doctor, date, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":  "failed to get consultations",
			"response": err.Error(),
		})
	}

	// If there are existing doctor holidays, update the record
	if len(existingHolidays) > 0 {
		updatedDoctorHoliday, err := repository.DoctorInactiveSchedule(doctor, modelDoctorHoliday)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":  "failed to mark doctor as inactive",
				"response": err.Error(),
			})
		}

		for _, consultation := range consultations {
			// Send notification to each patient
			CreateNotification(
				consultation.PatientID,
				"Dokter Membatalkan Konsultasi!",
				"Dokter membatalkan konsultasi karena urusan tertentu. Silakan daftar konsultasi di sesi atau hari lain.",
				"janti_temu",
			)
		}

		// Convert the updatedDoctorHoliday to DTO response
		doctorHolidayResponse := dto.ConvertToDoctorHolidayResponse(updatedDoctorHoliday)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success, doctor marked as inactive",
			"response": doctorHolidayResponse,
		})
	}

	// If there are no existing doctor holidays, create a new record
	doctorHoliday, err = repository.DoctorInactiveSchedule(doctor, modelDoctorHoliday)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":  "failed to create doctor holiday",
			"response": err.Error(),
		})
	}

	for _, consultation := range consultations {
		// Send notification to each patient
		CreateNotification(
			consultation.PatientID,
			"Dokter Membatalkan Konsultasi!",
			"Dokter membatalkan konsultasi karena urusan tertentu. Silakan daftar konsultasi di sesi atau hari lain.",
			"janti_temu",
		)
	}

	doctorHolidayResponse := dto.ConvertToDoctorHolidayResponse(doctorHoliday)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success, doctor marked as inactive",
		"response": doctorHolidayResponse,
	})
}
