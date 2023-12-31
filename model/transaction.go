package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionStatus string

const (
	Waiting   TransactionStatus = "waiting"
	Processed TransactionStatus = "processed"
	Done      TransactionStatus = "done"
	Cancelled TransactionStatus = "cancelled"
)

type Transaction struct {
	gorm.Model
	ID             uuid.UUID         `json:"id" form:"id"`
	ConsultationID uuid.UUID         `gorm:"index" json:"consultation_id"`
	Invoice        string            `gorm:"size:255"`
	Price          float64           `gorm:"type:decimal(15,2)"`
	AdminPrice     float64           `gorm:"type:decimal(15,2)"`
	Total          float64           `gorm:"type:decimal(15,2)"`
	Status         TransactionStatus `gorm:"type:ENUM('waiting', 'processed', 'done', 'cancelled')"`
	PaymentStatus  string            `gorm:"type:ENUM('pending', 'done', 'refund')"`
	Refund         Refund            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment        Payment           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consultation   Consultation      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
