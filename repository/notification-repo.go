package repository

import (
	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetAllNotifications(patient_id uuid.UUID, category string) ([]model.Notification, error) {
	var datanotifications []model.Notification

	tx := database.DB

	if category != "" {
        tx = tx.Where("category LIKE ?", "%"+category+"%")
    }

	tx.Where("patient_id = ?", patient_id).Find(&datanotifications)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datanotifications, nil
}

func CreateNotification(data model.Notification){
	database.DB.Save(&data)
}