package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type NotificationResponse struct {
	ID          uuid.UUID `json:"id"`
	PatientID   uuid.UUID `json:"patient_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

func ConvertToNotificationModel(patient_id uuid.UUID, title, content, category string) model.Notification {
	return model.Notification{
		ID:        uuid.New(),
		PatientID: patient_id,
		Title:     title,
		Content:   content,
		Category:  category,
		Date: 	   time.Now(),
	}
}

func ConvertToNotificationResponse(notification model.Notification) NotificationResponse {
	return NotificationResponse{
		ID:        notification.ID,
		PatientID: notification.PatientID,
		Title  :   notification.Title,
		Content:   notification.Content,
		Category:  notification.Category,
		Date: 	   notification.Date,
	}
}