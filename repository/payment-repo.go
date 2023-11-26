package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func InsertPayment(data model.Payment) (model.Payment, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Payment{}, tx.Error
	}
	
	return data, nil
}

func CheckPayment(id uuid.UUID) bool {
	var payment model.Payment

	if err := database.DB.Where("transaction_id = ?", id).First(&payment).Error; err != nil {
        return false
	}

	return true
}