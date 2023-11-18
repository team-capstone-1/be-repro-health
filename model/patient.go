package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID       		   uuid.UUID `json:"id" form:"id"`
	Name     		   string    `gorm:"size:255"`
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
