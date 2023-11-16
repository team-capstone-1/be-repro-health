package repository

import (
	"capstone-project/model"
	"gorm.io/gorm"
)

type DoctorArticleRepository struct {
	DB *gorm.DB
}

func NewDoctorArticleRepository(db *gorm.DB) *DoctorArticleRepository {
	return &DoctorArticleRepository{
		DB: db,
	}
}

func (r *DoctorArticleRepository) GetAllArticles() ([]model.Article, error) {
	var articles []model.Article
	if err := r.DB.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *DoctorArticleRepository) GetArticleByID(id uint) (*model.Article, error) {
	var article model.Article
	if err := r.DB.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *DoctorArticleRepository) CreateArticle(article *model.Article) error {
	tx := r.DB.Begin()
	if err := tx.Create(article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *DoctorArticleRepository) UpdateArticle(article *model.Article) error {
	tx := r.DB.Begin()
	if err := tx.Save(article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *DoctorArticleRepository) DeleteArticle(id uint) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&model.Article{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *DoctorArticleRepository) GetArticlesByDoctorID(doctorID uint) ([]model.Article, error) {
	var articles []model.Article
	if err := r.DB.Where("doctor_id = ?", doctorID).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}