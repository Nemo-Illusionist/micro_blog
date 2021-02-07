package version

import (
	"gorm.io/gorm"
	"micro_blog/config"
	"micro_blog/dal/models"
	"micro_blog/migration/contract"
)

type M202101252200Init struct {
}

func (u M202101252200Init) Name() string {
	return "202101252200_init"
}

func (u M202101252200Init) PreviousMigration() *contract.Migration {
	return nil
}

func (u M202101252200Init) Up(db *gorm.DB) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
	if err != nil {
		return err
	}

	user := &models.User{Name: "Admin", Login: "admin", IsAdmin: true}
	user.SetPasswordHash("admin", cfg.PasswordSecretKey)
	err = db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u M202101252200Init) Down(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Article{})
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&models.User{})
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&models.Comment{})
	if err != nil {
		return err
	}

	return nil
}
