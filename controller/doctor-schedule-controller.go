package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/repository"
	"net/http"

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
	responseData, err := repository.DoctorGetAllSchedules(doctor)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "failed get schedules data",
			"response": err.Error(),
		})
	}

	checkPatient, err := repository.GetPatientByID()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get patient",
			"reponse": err.Error(),
		})
	}

	var doctorSchedule []dto.DoctorGetSchedule
	for _, schedule := range responseData {
		doctorSchedule = append(doctorSchedule, dto.ConvertToDoctorScheduleResponse(schedule))
		var patientID = schedule.PatientID
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get schedule data",
		"response": doctorSchedule,
	})
}
