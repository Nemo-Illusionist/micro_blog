package migration

import "time"

type MigrateVersion struct {
	ID        uint32    `gorm:"primary_key;"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func (MigrateVersion) TableName() string {
	return "migrate_version"
}
