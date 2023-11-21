package repository

import (
	"capstone-project/database"
	"capstone-project/model"
)

func GetAllTransactions() ([]model.Transaction, error) {
	var datatransactions []model.Transaction

	tx := database.DB.Find(&datatransactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datatransactions, nil
}
