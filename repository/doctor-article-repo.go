package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetAllArticles(doctor_id string) ([]model.Article, error) {
	var dataarticles []model.Article

	tx := database.DB

	if doctor_id != "" {
		tx = tx.Where("doctor_id = ?", doctor_id)
	}

	tx.Find(&dataarticles)
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

func InsertArticle(data model.Article) (model.Article, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return data, nil
}

func DeleteArticleByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Article{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// type DoctorArticleRepository struct {
// 	DB *gorm.DB
// }

// func NewDoctorArticleRepository(db *gorm.DB) *DoctorArticleRepository {
// 	return &DoctorArticleRepository{
// 		DB: db,
// 	}
// }

// func (r *DoctorArticleRepository) GetAllArticles() ([]model.Article, error) {
// 	var articles []model.Article
// 	if err := r.DB.Find(&articles).Error; err != nil {
// 		return nil, err
// 	}
// 	return articles, nil
// }

// func (r *DoctorArticleRepository) GetArticleByID(id uint) (*model.Article, error) {
// 	var article model.Article
// 	if err := r.DB.First(&article, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &article, nil
// }

// func (r *DoctorArticleRepository) CreateArticle(article *model.Article) error {
// 	tx := r.DB.Begin()
// 	if err := tx.Create(article).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	return nil
// }

// func (r *DoctorArticleRepository) UpdateArticle(article *model.Article) error {
// 	tx := r.DB.Begin()
// 	if err := tx.Save(article).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	return nil
// }

// func (r *DoctorArticleRepository) DeleteArticle(id uint) error {
// 	tx := r.DB.Begin()
// 	if err := tx.Delete(&model.Article{}, id).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	return nil
// }

// func (r *DoctorArticleRepository) GetArticlesByDoctorID(doctorID uint) ([]model.Article, error) {
// 	var articles []model.Article
// 	if err := r.DB.Where("doctor_id = ?", doctorID).Find(&articles).Error; err != nil {
// 		return nil, err
// 	}
// 	return articles, nil
// }
