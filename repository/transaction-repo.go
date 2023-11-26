package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"

	"github.com/google/uuid"
)

func GetDoctorTransactions(doctorID uuid.UUID) ([]model.Transaction, error) {
	var transactions []model.Transaction

	tx := database.DB.Where("doctor_id = ?", doctorID).Preload("Consultation").Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return transactions, nil
}

func GetTransactionByID(id uuid.UUID) (model.Transaction, error) {
	var datatransaction model.Transaction

	tx := database.DB.Preload("Consultation").Preload("Consultation.Clinic").Preload("Consultation.Doctor").First(&datatransaction, id)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return datatransaction, nil
}

func InsertTransaction(data model.Transaction) (model.Transaction, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return data, nil
}

func GenerateNextInvoice() (string, time.Time, error) {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()

	formattedInvoice := fmt.Sprintf("INV/%d/%02d/%02d/", year, month, day)

	var lastInvoice model.Transaction
	if err := database.DB.Where("invoice LIKE ?", formattedInvoice+"%").Order("invoice DESC").First(&lastInvoice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			formattedInvoice += "0001"
		} else {
			return "", now, err
		}
	} else {
		var sequence int
		_, err := fmt.Sscanf(lastInvoice.Invoice, formattedInvoice+"%04d", &sequence)
		if err != nil {
			return "", now, err
		}

		formattedInvoice += fmt.Sprintf("%04d", sequence+1)
	}

	return formattedInvoice, now, nil
}

func GetTransactionByID(id uuid.UUID) (model.Transaction, error) {
	var datatransaction model.Transaction

	tx := database.DB.Preload("Consultation").Preload("Consultation.Clinic").Preload("Consultation.Doctor").First(&datatransaction, id)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return datatransaction, nil
}

func InsertTransaction(data model.Transaction) (model.Transaction, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return model.Transaction{}, tx.Error
	}
	return data, nil
}

func GenerateNextInvoice() (string, time.Time, error) {
    now := time.Now()
    year, month, day := now.Year(), now.Month(), now.Day()

    formattedInvoice := fmt.Sprintf("INV/%d/%02d/%02d/", year, month, day)

    var lastInvoice model.Transaction
    if err := database.DB.Where("invoice LIKE ?", formattedInvoice+"%").Order("invoice DESC").First(&lastInvoice).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            formattedInvoice += "0001"
        } else {
            return "", now, err
        }
    } else {
        var sequence int
        _, err := fmt.Sscanf(lastInvoice.Invoice, formattedInvoice+"%04d", &sequence)
        if err != nil {
            return "", now, err
        }

        formattedInvoice += fmt.Sprintf("%04d", sequence+1)
    }

    return formattedInvoice, now, nil
}