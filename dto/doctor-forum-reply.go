package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorForumReplyRequest struct {
	ForumsID uuid.UUID `json:"forum_id" form:"forum_id"`
	Content  string    `json:"content" form:"content"`
	Date     time.Time `json:"date" form:"date"`
}

type DoctorForumReplyResponse struct {
	ID       uuid.UUID `json:"id"`
	ForumsID uuid.UUID `json:"forum_id"`
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

func ConvertToDoctorReplyModel(forum DoctorForumReplyRequest) model.ForumReply {
	return model.ForumReply{
		ID:       uuid.New(),
		ForumsID: forum.ForumsID,
		Content:  forum.Content,
		Date:     time.Now(),
	}
}

func ConvertToDoctorForumReplyResponse(forum model.ForumReply) DoctorForumReplyResponse {
	return DoctorForumReplyResponse{
		ID:       forum.ID,
		ForumsID: forum.ForumsID,
		Content:  forum.Content,
		Date:     forum.Date,
	}
}
