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
	var date time.Time
	if dateString != "" {
		_, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message":  "failed to parse date",
				"response": err.Error(),
			})
		}
	}

	responseData, err := repository.DoctorGetAllSchedules(doctor, session, date)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "failed get schedules data",
			"response": err.Error(),
		})
	}

	var doctorSchedules []dto.DoctorGetSchedule
	for _, schedule := range responseData {
		patient, err := repository.GetPatientByID(schedule.PatientID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"message": "failed get patient",
				"reponse": err.Error(),
			})
		}
		doctorSchedule := dto.ConvertToDoctorScheduleResponse(schedule, patient)
		doctorSchedules = append(doctorSchedules, doctorSchedule)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get schedule data",
		"response": doctorSchedules,
	})
}
