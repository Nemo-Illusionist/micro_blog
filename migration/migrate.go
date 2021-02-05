package migration

import (
	"gorm.io/gorm"
	"micro_blog/migration/contract"
	"strings"
	"time"
)

func Migrate(db *gorm.DB, migrationsMap []contract.Migration, mod string) error {
	err := db.AutoMigrate(&MigrateVersion{})
	if err != nil {
		return err
	}

	var versions []MigrateVersion
	db.Order("id").Find(&versions)
	if strings.EqualFold(mod, Up) {
		for _, m := range migrationsMap[len(versions):] {
			t := db.Begin()
			err = db.Transaction(m.Up)
			if err != nil {
				return err
			}
			db.Create(&MigrateVersion{Name: m.Name(), CreatedAt: time.Now().UTC()})
			t.Commit()
		}
	} else if strings.EqualFold(mod, Down) {
		for i, m := range migrationsMap[:len(versions)] {
			t := db.Begin()
			err = db.Transaction(m.Down)
			if err != nil {
				return err
			}
			db.Where(&versions[i]).Delete(&versions[i])
			t.Commit()
		}
	}

	return nil
}
