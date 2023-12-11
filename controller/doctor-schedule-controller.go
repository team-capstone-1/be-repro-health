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

	if doctorHoliday.DoctorAvailable == false {
		patientIDs, err := repository.GetPatientIDsByDateAndSession(doctorID, dateString, session)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message":  "failed to get patient IDs",
				"response": err.Error(),
			})
		}

		for _, patientID := range patientIDs {
			CreateNotification(
				patientID,
				"Dokter Tidak Tersedia",
				"Dokter tidak tersedia pada sesi ini. Silakan cek jadwal dokter untuk sesi atau hari lain.",
				"janji_temu",
			)
		}
	}

	// Update transaction status to "waiting"
	err = repository.UpdateTransactionStatusToWaiting(dateString, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed to update transaction status",
			"response": err.Error(),
		})
	}

	responseData, err := repository.DoctorGetAllSchedules(doctorID, session, dateString)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "failed get schedules data",
			"response": err.Error(),
		})
	}

	doctorHolidayResponse := dto.ConvertToDoctorScheduleResponse(doctorID, responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success, doctor status updated",
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

	doctorHoliday, err := repository.DoctorActiveSchedule(doctorID, dateString, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed to mark doctor as active",
			"response": err.Error(),
		})
	}

	if doctorHoliday.DoctorAvailable == true {
		patientIDs, err := repository.GetPatientIDsByDateAndSession(doctorID, dateString, session)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message":  "failed to get patient IDs",
				"response": err.Error(),
			})
		}

		for _, patientID := range patientIDs {
			CreateNotification(
				patientID,
				"Dokter Ada!",
				"Dokter bersedia pada sesi ini. Silakan tunggu konfirmasi dari dokter yaa",
				"janji_temu",
			)
		}
	}

	responseData, err := repository.DoctorGetAllSchedules(doctorID, session, dateString)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message":  "failed get schedules data",
			"response": err.Error(),
		})
	}

	doctorHolidayResponse := dto.ConvertToDoctorScheduleResponse(doctorID, responseData)

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success, doctor status updated",
		"response": doctorHolidayResponse,
	})
}
