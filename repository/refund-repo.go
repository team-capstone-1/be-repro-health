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

func UpdateRefundStatus(id uuid.UUID) (model.Refund, error) {
	var datarefund model.Refund
    tx := database.DB.Model(&model.Refund{}).Where("id = ?", id).Update("status", "success")

    if tx.Error != nil {
        return model.Refund{}, tx.Error
    }

	tx = database.DB.First(&datarefund, id)
	if tx.Error != nil {
		return model.Refund{}, tx.Error
	}

	tx = database.DB.Model(&model.Transaction{}).Where("id = ?", datarefund.TransactionID).Update("status", "cancelled")
	if tx.Error != nil {
		return model.Refund{}, tx.Error
	}

    return datarefund, nil
}