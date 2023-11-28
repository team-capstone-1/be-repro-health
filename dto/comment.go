package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type CommentRequest struct {
	PatientID  uuid.UUID  `json:"patient_id" form:"patient_id"`
	Comment	   string     `json:"comment" form:"comment"`
}

type CommentResponse struct {
	ID    	  	   uuid.UUID  `json:"id"`
	ArticleID	   uuid.UUID  `json:"article_id"`
	PatientID  	   uuid.UUID  `json:"patient_id"`
	Comment		   string     `json:"comment"`
	Date  		   time.Time  `json:"date"`
}

func ConvertToCommentModel(comment CommentRequest) model.Comment {
	return model.Comment{
		ID:     	uuid.New(),
		PatientID:  comment.PatientID,
		Comment:    comment.Comment,
		Date: 		time.Now(),
	}
}

func ConvertToCommentResponse(comment model.Comment) CommentResponse {
	return CommentResponse{
		ID:    	   	comment.ID,
		ArticleID:  comment.ArticleID,
		PatientID:  comment.PatientID,
		Comment:    comment.Comment,
		Date: 		comment.Date,
	}
}
