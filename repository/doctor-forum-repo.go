package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func DoctorGetAllForums(title string, patient_id string) ([]model.Forum, error) {
	var dataforums []model.Forum

	tx := database.DB

	if title != "" {
		tx = tx.Where("title LIKE ?", "%"+title+"%")
	}

	if patient_id != "" {
		tx = tx.Where("patient_id = ?", patient_id)
	}

	tx.Find(&dataforums)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataforums, nil
}

func CreateDoctorReplyForum(data model.ForumReply) (model.ForumReply, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.ForumReply{}, tx.Error
	}
	return data, nil
}

func UpdateDoctorReplyForum(forumID uuid.UUID, data model.ForumReply) (model.ForumReply, error) {
	tx := database.DB.Where("id = ?", forumID).Updates(&data)
	if tx.Error != nil {
		return model.ForumReply{}, tx.Error
	}
	return data, nil
}

func GetDoctorReplyForumByID(forumID uuid.UUID) (model.ForumReply, error) {
	var data model.ForumReply
	tx := database.DB.Where("id = ?", forumID).First(&data)

	if tx.Error != nil {
		return model.ForumReply{}, tx.Error
	}

	return data, nil
}
