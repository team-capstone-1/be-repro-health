package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func InsertPayment(data model.Payment) (model.Payment, error) {
    tx := database.DB.Begin()

    if err := tx.Error; err != nil {
        return model.Payment{}, err
    }

    if err := tx.Save(&data).Error; err != nil {
        tx.Rollback()
        return model.Payment{}, err
    }

    if err := tx.Model(&model.Transaction{}).Where("id = ?", data.TransactionID).Updates(map[string]interface{}{"payment_status": "done"}).Error; err != nil {
        tx.Rollback()
        return model.Payment{}, err
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return model.Payment{}, err
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