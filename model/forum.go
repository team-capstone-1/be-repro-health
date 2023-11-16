package model

import (
	"time"

	"gorm.io/gorm"
)

type Forums struct {
	gorm.Model
	title      string     `gorm:"size:255"`
	content    string     `gorm:"size:255"`
	anonymous  bool       `gorm:"default:false"`
	date       time.Time  `gorm:"type:datetime"`
	ForumReply ForumReply `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
