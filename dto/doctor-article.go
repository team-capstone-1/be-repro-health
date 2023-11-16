package dto

import (
	"capstone-project/model"
	"time"
)

type DoctorArticleRequest struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
	Image   string    `json:"image"`
}

type DoctorArticleResponse struct {
	ID       uint      `json:"id"`
	DoctorID uint      `json:"doctor_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
	Image    string    `json:"image"`
}

func ConvertToDoctorArticleModel(doctor *DoctorArticleRequest, doctorID uint) *model.Article {
	return &model.Article{
		DoctorID: doctorID,
		Title:    doctor.Title,
		Content:  doctor.Content,
		Date:     time.Now(),
	}
}

func ConvertToDoctorArticleResponse(article *model.Article) *DoctorArticleResponse {
	return &DoctorArticleResponse{
		ID:       article.ID,
		DoctorID: article.DoctorID,
		Title:    article.Title,
		Content:  article.Content,
		Date:     article.Date,
		Image:    article.Image,
	}
}

// func ConvertToDoctorArticleResponseList(articles []*model.Article) []*DoctorArticleResponse {
// 	var responseList []*DoctorArticleResponse
// 	for _, article := range articles {
// 		response := ConvertToDoctorArticleResponse(article)
// 		responseList = append(responseList, response)
// 	}
// 	return responseList
// }
