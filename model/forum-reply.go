package model

import (
	"time"

	"gorm.io/gorm"
)

type ForumReply struct {
	gorm.Model
	ForumsID uint      `gorm:"index" json:"forums_id"`
	Content  string    `gorm:"size:255"`
	Date     time.Time `gorm:"type:datetime"`
}
