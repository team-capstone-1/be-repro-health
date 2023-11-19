package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetAllDoctorsArticles() ([]model.Article, error) {
	var datadoctorarticles []model.Article

	tx := database.DB.Find(&datadoctorarticles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datadoctorarticles, nil
}

func GetDoctorArticleByID(id uuid.UUID) (model.Article, error) {
	var datadoctorarticles model.Article

	tx := database.DB.First(&datadoctorarticles, id)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return datadoctorarticles, nil
}

func InsertDoctorArticle(data model.Article) (model.Article, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return data, nil
}

func DeleteDoctorArticleByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Article{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateDoctorArticleByID(id uuid.UUID, updateData model.Article) (model.Article, error) {
	tx := database.DB.Model(&updateData).Where("id = ?", id).Updates(updateData)
	if tx.Error != nil {
		return model.Article{}, tx.Error
	}
	return updateData, nil
}