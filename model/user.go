package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	Name     string    `json:"name" form:"name"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
	OTP 	 string    `json:"otp" form:"otp"`
	Patients []Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Bookmarks []Article `gorm:"many2many:article_bookmark"`
}
