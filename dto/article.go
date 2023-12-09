package dto

import (
	"capstone-project/model"
	"capstone-project/repository"
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
	Profile		   string     `json:"patient_profile"`
	Date  		   time.Time  `json:"date"`
}

type UserArticleResponse struct {
	ID        uuid.UUID         `json:"id"`
	DoctorID  uuid.UUID         `json:"doctor_id"`
	Title     string            `json:"title"`
	Tags      string            `json:"tags"`
	Reference string            `json:"reference"`
	Date      time.Time         `json:"date"`
	Image     string            `json:"image"`
	ImageDesc string            `json:"image_desc"`
	Content   string            `json:"content"`
	View      int               `json:"views"`
	Bookmark  bool 				`json:"bookmarked"`
	Profile   string			`json:"doctor_profile"`
	Comment   []CommentResponse `json:"comments"`
}

func ConvertToCommentModel(comment CommentRequest) model.Comment {
	return model.Comment{
		ID:     	uuid.New(),
		PatientID:  comment.PatientID,
		Comment:    comment.Comment,
		Date: 		time.Now(),
	}
}

func ConvertToCommentResponse(comment model.Comment, profile string) CommentResponse {
	return CommentResponse{
		ID:    	   	comment.ID,
		ArticleID:  comment.ArticleID,
		PatientID:  comment.PatientID,
		Comment:    comment.Comment,
		Date: 		comment.Date,
		Profile: 	profile,
	}
}

func ConvertToUserArticleResponse(article model.Article, bookmarked bool) UserArticleResponse {
	var articleCommentResponses []CommentResponse

	for _, reply := range article.Comment {
		articleCommentResponses = append(articleCommentResponses, ConvertToCommentResponse(reply, repository.GetProfileByPatientID(reply.PatientID)))
	}

	return UserArticleResponse{
		ID:        article.ID,
		DoctorID:  article.DoctorID,
		Title:     article.Title,
		Tags:      article.Tags,
		Reference: article.Reference,
		Date:      article.Date,
		Image:     article.Image,
		ImageDesc: article.ImageDesc,
		Content:   article.Content,
		View:      article.View,
		Bookmark:  bookmarked,
		Profile:   repository.GetProfileByDoctorID(article.DoctorID),
		Comment:   articleCommentResponses,
	}
}
