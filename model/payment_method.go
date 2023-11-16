package model

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	TransactionID uint   `gorm:"index" json:"transaction_id"`
	Name          string `gorm:"size:255"`
	AccountNumber string `gorm:"size:255"`
	Image         string `gorm:"size:255"`
}
