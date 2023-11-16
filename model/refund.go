package model

import "gorm.io/gorm"

type RefundStatus string

const (
	Processing RefundStatus = "processing"
	Success    RefundStatus = "success"
)

type Refunds struct {
	gorm.Model
	TransactionID uint         `gorm:"index" json:"transaction_id"`
	Name          string       `gorm:"size:255"`
	AccountNumber string       `gorm:"size:255"`
	Status        RefundStatus `gorm:"type:ENUM('processing', 'success')"`
}
