package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetSpecialists(id uuid.UUID) ([]model.Specialist, error) {
	var specialists []model.Specialist

	tx := database.DB.Where("id = ?", id).Find(&specialists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return specialists, nil
}

func GetSpecialistByID(id uuid.UUID) (model.Specialist, error) {
	var specialist model.Specialist

	tx := database.DB.First(&specialist, id)
	if tx.Error != nil {
		return specialist, tx.Error
	}

	return specialist, nil
}

func InsertSpecialist(data model.Specialist) (model.Specialist, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Specialist{}, tx.Error
	}
	return data, nil
}

func DeleteSpecialistByID(id uuid.UUID) error {
	tx := database.DB.Delete(&model.Specialist{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateSpecialist(data model.Specialist) (model.Specialist, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Specialist{}, tx.Error
	}
	return data, nil
}