package repository

import (
	"gorm.io/gorm"
	"fmt"
	"time"

	"capstone-project/database"
	"capstone-project/model"

	"github.com/google/uuid"
)

func GetAllTransactionsByDoctorID(doctorID uuid.UUID) ([]model.Transaction, error) {
	var datatransactions []model.Transaction

	tx := database.DB.Preload("Consultation").
		Joins("JOIN consultations ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ?", doctorID).
		Find(&datatransactions)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return datatransactions, nil
}

func GetDoneTransactionsByDoctorAndMonth(doctorID uuid.UUID, month time.Time) ([]model.Transaction, error) {
	var transactions []model.Transaction
	startOfMonth := month.AddDate(0, 0, 1)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	tx := database.DB.
		Preload("Consultation").
		Joins("JOIN consultations ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND transactions.date BETWEEN ? AND ? AND transactions.payment_status = 'done'", doctorID, startOfMonth, endOfMonth).
		Find(&transactions)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return transactions, nil
}

func GetDoneTransactionsByDoctorAndWeek(doctorID uuid.UUID, week time.Time) ([]model.Transaction, error) {
	var transactions []model.Transaction
	startOfWeek := week.AddDate(0, 0, -7)
	endOfWeek := startOfWeek.AddDate(0, 0, -14)

	tx := database.DB.
		Preload("Consultation").
		Joins("JOIN consultations ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND transactions.date BETWEEN ? AND ? AND transactions.payment_status = 'done'", doctorID, startOfWeek, endOfWeek).
		Find(&transactions)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return transactions, nil
}

func GetDoneTransactionsByDoctorAndTimeRange(doctorID uuid.UUID, start, end time.Time) ([]model.Transaction, error) {
	var transactions []model.Transaction

	tx := database.DB.
		Preload("Consultation").
		Joins("JOIN consultations ON transactions.consultation_id = consultations.id").
		Where("consultations.doctor_id = ? AND transactions.date BETWEEN ? AND ? AND transactions.payment_status = 'done'", doctorID, start, end).
		Find(&transactions)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return transactions, nil
}

func GetPatientTransactions(id uuid.UUID) ([]model.Transaction, error) {
	var datatransactions []model.Transaction

	tx := database.DB.
		Preload("Refund").Preload("Payment").Preload("Consultation").Preload("Consultation.Clinic").Preload("Consultation.Doctor").
		Joins("JOIN consultations ON transactions.consultation_id = consultations.id").
		Where("consultations.patient_id = ?", id).
		Find(&datatransactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datatransactions, nil
}

func GetTransactionByID(id uuid.UUID) (model.Transaction, error) {
	var datatransaction model.Transaction

	tx := database.DB.Preload("Refund").Preload("Payment").Preload("Consultation").Preload("Consultation.Clinic").Preload("Consultation.Doctor").First(&datatransaction, id)
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

func UpdateTransactionStatus(id uuid.UUID, status string) error {
	var datatransaction model.Transaction
    tx := database.DB.Model(&datatransaction).Where("id = ?", id).Update("status", status)

    if tx.Error != nil {
        return tx.Error
    }

    return nil
}

func UpdateTransactionPaymentStatus(id uuid.UUID, status string) error {
	var datatransaction model.Transaction
    tx := database.DB.Model(&datatransaction).Where("id = ?", id).Update("payment_status", status)

    if tx.Error != nil {
        return tx.Error
    }

    return nil
}