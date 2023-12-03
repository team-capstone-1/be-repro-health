package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorArticleRequest struct {
	Title     string `json:"title" form:"title"`
	Tags      string `json:"tags" form:"tags"`
	Reference string `json:"reference" form:"reference"`
	Image     string `json:"image" form:"image"`
	ImageDesc string `json:"image_desc" form:"image_desc"`
	Content   string `json:"content" form:"content"`
}

type DoctorArticleResponse struct {
	ID        uuid.UUID         `json:"id"`
	DoctorID  uuid.UUID         `json:"doctor_id"`
	Title     string            `json:"title"`
	Tags      string            `json:"tags"`
	Reference string            `json:"reference"`
	Date      time.Time         `json:"date"`
	Image     string            `json:"image"`
	ImageDesc string            `json:"image_desc"`
	Content   string            `json:"content"`
	Published bool              `json:"published"`
	View      int               `json:"views"`
	Comment   []CommentResponse `json:"comments"`
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
	Comment   []CommentResponse `json:"comments"`
}

func ConvertToDoctorArticleModel(doctor DoctorArticleRequest) model.Article {
	return model.Article{
		ID:        uuid.New(),
		Title:     doctor.Title,
		Tags:      doctor.Tags,
		Reference: doctor.Reference,
		Image:     doctor.Image,
		ImageDesc: doctor.ImageDesc,
		Content:   doctor.Content,
		Date:      time.Now(),
		Published: false,
		View:      0,
		Comment:   []model.Comment{},
	}
}

func ConvertToDoctorArticleResponse(article model.Article) DoctorArticleResponse {
	return DoctorArticleResponse{
		ID:        article.ID,
		DoctorID:  article.DoctorID,
		Title:     article.Title,
		Tags:      article.Tags,
		Reference: article.Reference,
		Date:      article.Date,
		Image:     article.Image,
		ImageDesc: article.ImageDesc,
		Content:   article.Content,
		Published: article.Published,
		View:      article.View,
		Comment:   []CommentResponse{},
	}
}

func ConvertToDoctorArticleDashboardResponse(article model.Article) DoctorArticleResponse {
	return DoctorArticleResponse{
		ID:        article.ID,
		DoctorID:  article.DoctorID,
		Title:     article.Title,
		Tags:      article.Tags,
		Reference: article.Reference,
		Date:      article.Date,
		Image:     article.Image,
		ImageDesc: article.ImageDesc,
		Content:   article.Content,
	}
}

func ConvertToUserArticleResponse(article model.Article) UserArticleResponse {
	var articleCommentResponses []CommentResponse

	for _, reply := range article.Comment {
		articleCommentResponses = append(articleCommentResponses, ConvertToCommentResponse(reply))
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
		Comment:   articleCommentResponses,
	}
}
