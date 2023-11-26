package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func DoctorGetAllForums(title string, patientID string, forumID uuid.UUID) ([]model.Forum, error) {
	var dataForums []model.Forum

	tx := database.DB.Model(model.Forum{}).Preload("ForumReply")

	if title != "" {
		tx = tx.Where("title LIKE ?", "%"+title+"%")
	}

	if patientID != "" {
		tx = tx.Where("patient_id = ?", patientID)
	}

	if err := tx.Find(&dataForums).Error; err != nil {
		return nil, err
	}

	if forumID != uuid.Nil {
		for i := range dataForums {
			var forumReplies []model.ForumReply
			if err := database.DB.Where("forums_id = ?", dataForums[i].ID).First(&forumReplies).Error; err != nil {
				return nil, err
			}

			if !dataForums[i].Status {
				dataForums[i].ForumReply = nil
			} else {
				dataForums[i].ForumReply = forumReplies
			}
		}
	}

	return dataForums, nil
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

func GetDoctorForumReplyByID(forumID uuid.UUID) (model.ForumReply, error) {
	var data model.ForumReply
	tx := database.DB.Where("id = ?", forumID).First(&data)

	if tx.Error != nil {
		return model.ForumReply{}, tx.Error
	}

	return data, nil
}

func DeleteForumReplyByID(id uuid.UUID) error {
	var forumReply []model.ForumReply
	tx := database.DB.Delete(&forumReply, "id = ?", id) // SOFT DELETE
	// tx := database.DB.Unscoped().Delete(&forumReply, "id = ?", id) // HARD DELETE
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateForumStatus(data model.Forum, status bool) error {
	tx := database.DB.Model(&data).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
