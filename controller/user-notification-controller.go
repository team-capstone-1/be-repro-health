package controller

import (
	"net/http"

	"capstone-project/repository"
	"capstone-project/dto"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

func CreateNotification(patient_id uuid.UUID, title, content, category string){
	notificationModel := dto.ConvertToNotificationModel(patient_id, title, content, category)
	repository.CreateNotification(notificationModel)
}

func GetNotificationsController(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}
	category := c.FormValue("category")

	responseData, err := repository.GetAllNotifications(uuid, category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get notifications",
			"response":   err.Error(),
		})
	}

	var notificationResponse []dto.NotificationResponse
	for _, notification := range responseData {
		notificationResponse = append(notificationResponse, dto.ConvertToNotificationResponse(notification))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get notifications",
		"response":   notificationResponse,
	})
}