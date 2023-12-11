package repository

import (
	"errors"

	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func InsertComment(data model.Comment) (model.Comment, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Comment{}, tx.Error
	}
	return data, nil
}

func UpdateBookmark(userID, articleID uuid.UUID) error {
	if err := database.DB.First(&model.User{}, "id = ?", userID).Error; err != nil {
		return errors.New("User ID not found")
	}

	if err := database.DB.First(&model.Article{}, "id = ?", articleID).Error; err != nil {
		return errors.New("Article not found")
	}

	var user model.User
	if err := database.DB.Preload("Bookmarks").First(&user, "id = ?", userID).Error; err != nil {
		return errors.New("Error loading user with bookmarks")
	}

	bookmarked := false
	for _, bookmark := range user.Bookmarks {
		if bookmark.ID == articleID {
			bookmarked = true
			break
		}
	}

	if bookmarked {
		if err := database.DB.Exec("DELETE FROM article_bookmark WHERE user_id = ? AND article_id = ?", userID, articleID).Error; err != nil {
			return err
		}
	} else {
		if err := database.DB.Exec("INSERT INTO article_bookmark (user_id, article_id) VALUES (?, ?)", userID, articleID).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetArticleBookmark(userID uuid.UUID) ([]model.Article, error){
	userData := model.User{}

	tx := database.DB.Preload("Bookmarks").Preload("Bookmarks.Comment").First(&userData, userID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return userData.Bookmarks, nil
}