package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethods struct {
	gorm.Model
	ID            uuid.UUID `json:"id" form:"id"`
	TransactionID uuid.UUID `gorm:"index" json:"transaction_id"`
	Name          string    `gorm:"size:255"`
	AccountNumber string    `gorm:"size:255"`
	Image         string    `gorm:"size:255"`
}
