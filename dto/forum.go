package dto

import (
	"time"

	"capstone-project/model"

	"github.com/google/uuid"
)

type ForumRequest struct {
	Title      string     `json:"title" form:"title"`
	Content    string     `json:"content" form:"content"`
	Anonymous  bool       `json:"anonymous" form:"anonymous"`
}

type ForumResponse struct {
	ID    uuid.UUID `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Anonymous  bool       `json:"anonymous"`
	Date       time.Time  `json:"date"`
}

func ConvertToForumModel(forum ForumRequest) model.Forum {
	return model.Forum{
		ID:       uuid.New(),
		Title:    forum.Title,
		Content:  forum.Content,
		Anonymous: forum.Anonymous,
		Date:     time.Now(),
	}
}

func ConvertToForumResponse(forum model.Forum) ForumResponse {
	return ForumResponse{
		ID:    	   forum.ID,
		Title:     forum.Title,
		Content:   forum.Content,
		Anonymous: forum.Anonymous,
		Date:      forum.Date,
	}
}
