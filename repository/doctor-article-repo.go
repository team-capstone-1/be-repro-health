package repository

import (
	"capstone-project/database"
	"capstone-project/model"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ...

func DoctorGetAllArticles(doctorID uuid.UUID) ([]model.Article, error) {
	var dataarticles []model.Article

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&dataarticles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataarticles, nil
}

func DoctorGetAllArticlesByMonth(doctorID uuid.UUID, month time.Time) ([]model.Article, error) {
	var dataarticles []model.Article

	startOfMonth := month.AddDate(0, 0, 1)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	tx := database.DB.Where("doctor_id = ? AND date BETWEEN ? AND ?", doctorID, startOfMonth, endOfMonth).Find(&dataarticles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataarticles, nil
}

func DoctorGetAllArticlesByWeek(doctorID uuid.UUID, week time.Time) ([]model.Article, error) {
	var dataarticles []model.Article

	startOfWeek := week.AddDate(0, 0, 0)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	tx := database.DB.Where("doctor_id = ? AND date BETWEEN ? AND ?", doctorID, startOfWeek, endOfWeek).Find(&dataarticles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataarticles, nil
}

func GetArticleByID(id uuid.UUID) (model.Article, error) {
	var dataarticle model.Article

	tx := database.DB.First(&dataarticle, id)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return dataarticle, nil
}

func UserGetAllArticle() ([]model.Article, error) {
	var dataarticlesdashboard []model.Article

	tx := database.DB.Find(&dataarticlesdashboard)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataarticlesdashboard, nil
}

func UserGetArticleByID(id uuid.UUID) (model.Article, error) {
	var dataarticle model.Article

	tx := database.DB.Preload("Comment").First(&dataarticle, id)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	database.DB.Model(&dataarticle).Where("id = ?", id).Updates(map[string]interface{}{"View": dataarticle.View + 1})

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
