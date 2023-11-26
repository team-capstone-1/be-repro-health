package dto

import (
	"time"

	"capstone-project/model"

	"github.com/google/uuid"
)

type ForumRequest struct {
	PatientID  uuid.UUID  `json:"patient_id" form:"patient_id"`
	Title      string     `json:"title" form:"title"`
	Content    string     `json:"content" form:"content"`
	Anonymous  bool       `json:"anonymous" form:"anonymous"`
}

type ForumResponse struct {
	ID    	   uuid.UUID `json:"id"`
	PatientID  uuid.UUID `json:"patient_id"`
	Title      string    `json:"title"`
	View       int       `json:"view"`
	Content    string    `json:"content"`
	Anonymous  bool      `json:"anonymous"`
	Date       time.Time `json:"date"`
}

func ConvertToForumModel(forum ForumRequest) model.Forum {
	return model.Forum{
		ID:       uuid.New(),
		PatientID:forum.PatientID,
		Title:    forum.Title,
		Content:  forum.Content,
		Anonymous: forum.Anonymous,
		Date:     time.Now(),
	}
}

func ConvertToForumResponse(forum model.Forum) ForumResponse {
	return ForumResponse{
		ID:    	   forum.ID,
		PatientID: forum.PatientID,
		Title:     forum.Title,
		View:      forum.View,
		Content:   forum.Content,
		Anonymous: forum.Anonymous,
		Date:      forum.Date,
	}
}
