package dto

import (
	"capstone-project/model"
	"time"

	"github.com/google/uuid"
)

type DoctorForumResponse struct {
	ID         uuid.UUID                  `json:"id"`
	PatientID  uuid.UUID                  `json:"patient_id"`
	Patient    PatientResponse            `json:"patient"`
	Title      string                     `json:"title"`
	Content    string                     `json:"content"`
	Anonymous  bool                       `json:"anonymous"`
	View       int                        `json:"view"`
	Date       time.Time                  `json:"date"`
	Status     bool                       `json:"status"`
	ForumReply []DoctorForumReplyResponse `json:"forum_replies"`
}

type ForumDoctorResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Price        float64   `json:"price"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profile_image"`
}

type DoctorForumReplyRequest struct {
	ForumsID uuid.UUID `json:"forum_id" form:"forum_id"`
	DoctorID uuid.UUID `json:"doctor_id" form:"doctor_id"`
	Content  string    `json:"content" form:"content"`
	Date     time.Time `json:"date" form:"date"`
}

type DoctorForumReplyResponse struct {
	ID       uuid.UUID           `json:"id"`
	ForumsID uuid.UUID           `json:"forums_id"`
	DoctorID uuid.UUID           `json:"doctor_id"`
	Doctor   ForumDoctorResponse `json:"doctor"`
	Content  string              `json:"content"`
	Date     time.Time           `json:"date"`
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

func ConvertToForumDoctorResponse(doctor model.Doctor) ForumDoctorResponse {
	return ForumDoctorResponse{
		ID:           doctor.ID,
		Name:         doctor.Name,
		Email:        doctor.Email,
		Price:        doctor.Price,
		Address:      doctor.Address,
		Phone:        doctor.Phone,
		ProfileImage: doctor.ProfileImage,
	}
}

func ConvertToDoctorForumReplyResponse(forumReply model.ForumReply) DoctorForumReplyResponse {
	return DoctorForumReplyResponse{
		ID:       forumReply.ID,
		ForumsID: forumReply.ForumsID,
		DoctorID: forumReply.DoctorID,
		Doctor:   ConvertToForumDoctorResponse(forumReply.Doctor),
		Content:  forumReply.Content,
		Date:     forumReply.Date,
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
		Patient:    ConvertToPatientResponse(forum.Patient),
		Title:      forum.Title,
		Content:    forum.Content,
		Anonymous:  forum.Anonymous,
		View:       forum.View,
		Date:       forum.Date,
		Status:     forum.Status,
		ForumReply: forumReplyResponses,
	}
}
