package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
	Patients []Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
