package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func InsertRefund(data model.Refund) (model.Refund, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Refund{}, tx.Error
	}
	return data, nil
}

func CheckRefund(id uuid.UUID) bool {
	var refund model.Refund

	if err := database.DB.Where("transaction_id = ?", id).First(&refund).Error; err != nil {
        return false
	}

	return true
}