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
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	session := c.FormValue("session")
	dateString := c.FormValue("date")

	if dateString != "" {
		_, err := time.Parse("02-01-2006", dateString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message":  "failed to parse date",
				"response": err.Error(),
			})
		}
	}

	responseData, err := repository.DoctorGetAllSchedules(doctor, session, dateString)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message":  "failed get schedules data",
			"response": err.Error(),
		})
	}

	doctorSchedules := dto.ConvertToDoctorScheduleResponse(doctor, responseData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get schedule data",
		"response": doctorSchedules,
	})
}

// func GetAllDoctorScheduleController(c echo.Context) error {
// 	doctorID := m.ExtractTokenUserId(c)
// 	if doctorID == uuid.Nil {
// 		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 			"message":  "unauthorized",
// 			"response": "Permission Denied: Doctor is not valid.",
// 		})
// 	}

// 	session := c.FormValue("session")
// 	dateString := c.FormValue("date")

// 	if dateString != "" {
// 		_, err := time.Parse("02-01-2006", dateString)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 				"message":  "failed to parse date",
// 				"response": err.Error(),
// 			})
// 		}
// 	}

// 	responseData, err := repository.DoctorGetAllSchedules(doctorID, session, dateString)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"message":  "failed get schedules data",
// 			"response": err.Error(),
// 		})
// 	}

// 	// Group schedules by date using a map
// 	scheduleMap := make(map[string][]model.Consultation)
// 	for _, schedule := range responseData {
// 		date := schedule.Date.Format("02-01-2006")
// 		scheduleMap[date] = append(scheduleMap[date], schedule)
// 	}

// 	// Convert the map to a slice of DoctorScheduleResponse
// 	var doctorSchedules []dto.DoctorScheduleResponse
// 	for _, schedules := range scheduleMap {
// 		doctorSchedule := dto.ConvertToDoctorScheduleResponse(doctorID, schedules)
// 		doctorSchedules = append(doctorSchedules, doctorSchedule)
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message":  "success get schedule data",
// 		"response": doctorSchedules,
// 	})
// }
