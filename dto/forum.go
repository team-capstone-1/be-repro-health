package dto

import (
	"time"

	"capstone-project/model"

	"github.com/google/uuid"
)

type ForumRequest struct {
	PatientID uuid.UUID `json:"patient_id" form:"patient_id"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	Anonymous bool      `json:"anonymous" form:"anonymous"`
}

type ForumResponse struct {
	ID        uuid.UUID `json:"id"`
	PatientID uuid.UUID `json:"patient_id"`
	Title     string    `json:"title"`
	View      int       `json:"view"`
	Content   string    `json:"content"`
	Profile   string    `json:"patient_profile"`
	Anonymous bool      `json:"anonymous"`
	Date      time.Time `json:"date"`
	Status    bool      `json:"status"`
	ForumReply []DoctorForumReplyResponse `json:"forum_replies"`
}

func ConvertToForumModel(forum ForumRequest) model.Forum {
	return model.Forum{
		ID:        uuid.New(),
		PatientID: forum.PatientID,
		Title:     forum.Title,
		Content:   forum.Content,
		Anonymous: forum.Anonymous,
		Date:      time.Now(),
	}
}

func ConvertToForumResponse(forum model.Forum) ForumResponse {
	var forumReplyResponses []DoctorForumReplyResponse

	for _, reply := range forum.ForumReply {
		forumReplyResponses = append(forumReplyResponses, ConvertToDoctorForumReplyResponse(reply))
	}

	return ForumResponse{
		ID:        forum.ID,
		PatientID: forum.PatientID,
		Title:     forum.Title,
		View:      forum.View,
		Content:   forum.Content,
		Anonymous: forum.Anonymous,
		Date:      forum.Date,
		Status:    forum.Status,
		ForumReply:forumReplyResponses,
	}
}
