package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorForumResponse struct {
	ID         uuid.UUID                  `json:"id"`
	PatientID  uuid.UUID                  `json:"patient_id"`
	Title      string                     `json:"title"`
	Content    string                     `json:"content"`
	Anonymous  bool                       `json:"anonymous"`
	Date       time.Time                  `json:"date"`
	Status     bool                       `json:"status"`
	ForumReply []DoctorForumReplyResponse `json:"forum_replies"`
}

type DoctorForumReplyRequest struct {
	ForumsID uuid.UUID `json:"forum_id" form:"forum_id"`
	DoctorID uuid.UUID `json:"doctor_id" form:"doctor_id"`
	Content  string    `json:"content" form:"content"`
	Date     time.Time `json:"date" form:"date"`
}

type DoctorForumReplyResponse struct {
	ID       uuid.UUID `json:"id"`
	ForumsID uuid.UUID `json:"forums_id"`
	DoctorID uuid.UUID `json:"doctor_id"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
}

type DoctorUpdateForumReplyRequest struct {
	Content string `json:"content" form:"content"`
}

func ConvertToDoctorUpdateForumReplyModel(forum DoctorUpdateForumReplyRequest) model.ForumReply {
	return model.ForumReply{
		Content: forum.Content,
	}
}

func ConvertToDoctorForumReplyModel(forum DoctorForumReplyRequest) model.ForumReply {
	return model.ForumReply{
		ID:       uuid.New(),
		ForumsID: forum.ForumsID,
		DoctorID: forum.DoctorID,
		Content:  forum.Content,
		Date:     time.Now(),
	}
}

func ConvertToDoctorForumReplyResponse(forum model.ForumReply) DoctorForumReplyResponse {
	return DoctorForumReplyResponse{
		ID:       forum.ID,
		ForumsID: forum.ForumsID,
		DoctorID: forum.DoctorID,
		Content:  forum.Content,
		Date:     forum.Date,
	}
}

func ConvertToDoctorForumResponse(forum model.Forum) DoctorForumResponse {
	var forumReplyResponses []DoctorForumReplyResponse

	for _, reply := range forum.ForumReply {
		forumReplyResponses = append(forumReplyResponses, ConvertToDoctorForumReplyResponse(reply))
	}

	return DoctorForumResponse{
		ID:         forum.ID,
		PatientID:  forum.PatientID,
		Title:      forum.Title,
		Content:    forum.Content,
		Anonymous:  forum.Anonymous,
		Date:       forum.Date,
		Status:     forum.Status,
		ForumReply: forumReplyResponses,
	}
}