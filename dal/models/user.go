package models

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Login     string    `gorm:"index:unique,where:deleted_at is null;not null"`
	Hash      string    `gorm:"not null"`
	Salt      string    `gorm:"not null"`
	IsAdmin   bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time
}

func (User) TableName() string {
	return "user"
}

func (u *User) CheckPassword(pass, secret string) bool {
	hash := _passToHash(pass, u.Salt, secret)
	return hash == u.Hash
}

func (u *User) SetPasswordHash(pass, secret string) {
	u.Salt = uuid.New().String()
	u.Hash = _passToHash(pass, u.Salt, secret)
}

func _passToHash(pass, salt, secret string) string {
	hash := sha256.New()
	hash.Write([]byte(pass + salt + secret))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
