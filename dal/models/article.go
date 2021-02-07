package models

import "time"

type Article struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	ShortBody string `gorm:"not null"`
	Body      string `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time

	UserID uint64 `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
}

func (Article) TableName() string {
	return "article"
}
