package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func InsertComment(data model.Comment) (model.Comment, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Comment{}, tx.Error
	}
	return data, nil
}