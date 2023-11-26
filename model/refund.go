package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RefundStatus string

const (
	Processing RefundStatus = "processing"
	Success    RefundStatus = "success"
)

type Refund struct {
	gorm.Model
	ID            uuid.UUID    `json:"id" form:"id"`
	TransactionID uuid.UUID    `gorm:"index" json:"transaction_id"`
	Name          string       `gorm:"size:255"`
	Date          time.Time    `gorm:"type:datetime"`
	AccountNumber string       `gorm:"size:255"`
	Status        RefundStatus `gorm:"type:ENUM('processing', 'success')"`
}
