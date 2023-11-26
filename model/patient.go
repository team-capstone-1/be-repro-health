package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID       		   uuid.UUID `json:"id" form:"id"`
	UserID 			   uuid.UUID `gorm:"column:user_id;index" json:"user_id"`
	User   			   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name     		   string    `gorm:"size:255"`
	TelephoneNumber    string    `gorm:"size:255"`
	ProfileImage       string    `gorm:"size:255"`
	DateOfBirth 	   time.Time `gorm:"type:date"`
	Relation		   string    `gorm:"size:255"`
	Weight			   float64   `gorm:"type:decimal(5,2)"`
	Height			   float64   `gorm:"type:decimal(5,2)"`
	KTPImage		   string    `gorm:"size:255"`
	NIK				   string    `gorm:"type:char(16)"`
	NoKartuKeluarga    string    `gorm:"type:char(16)"`
	KartuKeluargaImage string    `gorm:"size:255"`
}
