package models

import (
	"time"
	"gorm.io/datatypes"
)

type Note struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"index" json:"user_id"` // foreign key
	Title      string         `gorm:"type:varchar(255)" json:"title"`
	Content    string         `gorm:"type:text" json:"content"`
	Tags       datatypes.JSON `gorm:"type:json" json:"tags"` // JSON array olarak saklanır
	IsPinned   bool           `json:"is_pinned"`
	IsArchived bool           `json:"is_archived"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
