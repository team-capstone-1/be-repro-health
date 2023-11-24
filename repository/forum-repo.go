package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetAllForums(title string, patient_id string) ([]model.Forum, error) {
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

func GetForumByID(id uuid.UUID) (model.Forum, error) {
	var dataforum model.Forum

	tx := database.DB.First(&dataforum, id)
	if tx.Error != nil {
		return model.Forum{}, tx.Error
	}
	database.DB.Model(&dataforum).Where("id = ?", id).Updates(map[string]interface{}{"View": dataforum.View+1})

	return dataforum, nil
}

func InsertForum(data model.Forum) (model.Forum, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Forum{}, tx.Error
	}
	return data, nil
}

func DeleteForumByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Forum{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}