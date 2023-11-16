package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	Name     string    `json:"name" form:"name"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
}
