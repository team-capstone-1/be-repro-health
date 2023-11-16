package model

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}