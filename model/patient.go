package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DateOnly string

func (d *DateOnly) Scan(value interface{}) error {
	*d = DateOnly(value.(string))
	return nil
}

func (d DateOnly) Value() (string, error) {
	return string(d), nil
}

type Patient struct {
	gorm.Model
	ID       		   uuid.UUID `json:"id" form:"id"`
	UserID 			   uuid.UUID `gorm:"column:user_id;index" json:"user_id"`
	User   			   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name     		   string    `gorm:"size:255"`
	TelephoneNumber    string    `gorm:"size:255"`
	ProfileImage       string    `gorm:"size:255"`
	DateOfBirth 	   DateOnly  `gorm:"type:date"`
	Relation		   string    `gorm:"size:255"`
	Weight			   float64   `gorm:"type:decimal(5,2)"`
	Height			   float64   `gorm:"type:decimal(5,2)"`
	Gender             string    `gorm:"type:enum('male','female');default:'male'"`
}
