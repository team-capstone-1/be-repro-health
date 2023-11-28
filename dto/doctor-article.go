package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorArticleRequest struct {
	Title     string `json:"title"`
	Tags      string `json:"tags"`
	Reference string `json:"reference"`
	Image     string `json:"image"`
	ImageDesc string `json:"image_desc"`
	Content   string `json:"content"`
}

type DoctorArticleResponse struct {
	ID        uuid.UUID `json:"id"`
	DoctorID  uuid.UUID `json:"doctor_id"`
	Title     string    `json:"title"`
	Tags      string    `json:"tags"`
	Reference string    `json:"reference"`
	Date      time.Time `json:"date"`
	Image     string    `json:"image"`
	ImageDesc string    `json:"image_desc"`
	Content   string    `json:"content"`
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
