package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DoctorGetAllForums(title string, forumID uuid.UUID) ([]model.Forum, error) {
	var dataForums []model.Forum

	tx := database.DB.Model(model.Forum{}).Preload("ForumReply").Preload("Patient").Preload("ForumReply.Doctor")

	if title != "" {
		tx = tx.Where("title LIKE ?", "%"+title+"%")
	}

	if err := tx.Find(&dataForums).Error; err != nil {
		return nil, err
	}

	if forumID != uuid.Nil {
		var forumReplies []model.ForumReply

		for i := range dataForums {
			if err := database.DB.Where("forums_id = ?", dataForums[0].ID).First(&forumReplies).Error; err != nil {
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

func GetDoctorForumDetails(forumID uuid.UUID) ([]model.Forum, error) {
	var forum []model.Forum
	tx := database.DB.Where("id = ?", forumID).Preload("ForumReply").Preload("Patient").Preload("ForumReply.Doctor").First(&forum)

	if tx.Error != nil {
		return forum, tx.Error
	}

	if err := database.DB.Model(model.Forum{}).Where("id = ?", forumID).Update("View", gorm.Expr("view + ?", 1)).Error; err != nil {
		return forum, err
	}

	return forum, nil
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
