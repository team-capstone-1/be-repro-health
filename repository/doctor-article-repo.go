package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ...

func GetAllArticles(doctorID uuid.UUID) ([]model.Article, error) {
	var dataarticles []model.Article

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&dataarticles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataarticles, nil
}

func GetAllArticleDashboard() ([]model.Article, error) {
	var dataarticlesdashboard []model.Article

	tx := database.DB.Find(&dataarticlesdashboard)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataarticlesdashboard, nil
}

func GetArticleByID(id uuid.UUID) (model.Article, error) {
	var dataarticle model.Article

	tx := database.DB.First(&dataarticle, id)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return dataarticle, nil
}

func UserGetArticleByID(id uuid.UUID) (model.Article, error) {
	var dataarticle model.Article

	tx := database.DB.Preload("Comment").First(&dataarticle, id)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	database.DB.Model(&dataarticle).Where("id = ?", id).Updates(map[string]interface{}{"View": dataarticle.View+1})

	return dataarticle, nil
}

func InsertArticle(article model.Article) (model.Article, error) {
	tx := database.DB.Save(&article)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return article, nil
}

func UpdateArticle(article model.Article) error {
	result := database.DB.Save(&article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateArticlePublishedStatus(articleID uuid.UUID, doctorID uuid.UUID) error {
	article := model.Article{}

	// Cek apakah artikel ada dan dimiliki oleh dokter yang meminta update
	result := database.DB.Where("id = ? AND doctor_id = ?", articleID, doctorID).First(&article)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("article not found")
	} else if result.Error != nil {
		return result.Error
	}

	// Update status Published menjadi true
	article.Published = true

	// Simpan perubahan ke dalam database
	err := UpdateArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func DeleteArticleByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Article{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
