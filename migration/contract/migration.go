package contract

import "gorm.io/gorm"

type Migration interface {
	PreviousMigration() *Migration
	Name() string
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}
