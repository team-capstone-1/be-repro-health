package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorArticleRequest struct {
	DoctorID uuid.UUID `json:"doctor_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Image    string    `json:"image"`
}

type DoctorArticleResponse struct {
	ID       uuid.UUID `json:"id"`
	DoctorID uuid.UUID `json:"doctor_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
	Image    string    `json:"image"`
}

func ConvertToDoctorArticleModel(doctor DoctorArticleRequest) model.Article {
	return model.Article{
		ID:       uuid.New(),
		DoctorID: doctor.DoctorID,
		Title:    doctor.Title,
		Content:  doctor.Content,
		Date:     time.Now(),
	}
}

func ConvertToDoctorArticleResponse(article model.Article) DoctorArticleResponse {
	return DoctorArticleResponse{
		ID:       article.ID,
		DoctorID: article.DoctorID,
		Title:    article.Title,
		Content:  article.Content,
		Date:     article.Date,
		Image:    article.Image,
	}
}

func ConvertToDoctorArticleDashboardResponse(article model.Article) DoctorArticleResponse {
	return DoctorArticleResponse{
		ID:       article.ID,
		DoctorID: article.DoctorID,
	}
}