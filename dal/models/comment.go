package models

import "time"

type Comment struct {
	ID        uint64    `gorm:"primaryKey"`
	Body      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`

	ArticleID uint64  `gorm:"not null"`
	Article   Article `gorm:"foreignKey:ArticleID"`

	UserID uint64 `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`

	CommentID *uint64
	Comment   *Comment `gorm:"foreignKey:CommentID"`
}

func (Comment) TableName() string {
	return "comment"
}
