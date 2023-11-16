package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ForumReply struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	ForumsID uuid.UUID `gorm:"index" json:"forums_id"`
	Content  string    `gorm:"size:255"`
	Date     time.Time `gorm:"type:datetime"`
}
