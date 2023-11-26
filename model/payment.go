package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID            uuid.UUID `json:"id" form:"id"`
	TransactionID uuid.UUID `gorm:"index" json:"transaction_id"`
	Method        string    `gorm:"type:ENUM('manual_transfer', 'clinic_payment')"`
	Name          string    `gorm:"size:255"`
	AccountNumber string    `gorm:"size:255"`
	Image         string    `gorm:"size:255"`
}
