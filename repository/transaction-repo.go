package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetDoctorTransactions(doctorID uuid.UUID) ([]model.Transaction, error) {
	var transactions []model.Transaction

	tx := database.DB.Where("doctor_id = ?", doctorID).Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return transactions, nil
}
