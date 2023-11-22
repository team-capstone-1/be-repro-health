package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorForumResponse struct {
	ID        uuid.UUID `json:"id"`
	PatientID uuid.UUID `json:"patient_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Anonymous bool      `json:"anonymous"`
	Date      time.Time `json:"date"`
}

func ConvertToDoctorForumResponse(forum model.Forum) DoctorForumResponse {
	return DoctorForumResponse{
		ID:        forum.ID,
		PatientID: forum.PatientID,
		Title:     forum.Title,
		Content:   forum.Content,
		Anonymous: forum.Anonymous,
		Date:      forum.Date,
	}
}
