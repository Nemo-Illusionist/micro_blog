package dal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"micro_blog/config"
	"time"
)

func OpenWithConfig(conf *config.Config) (*gorm.DB, error) {
	ormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.Default.LogMode(0),
	}
	pgConfig := postgres.Config{
		DSN:                  conf.ConnectionString,
		PreferSimpleProtocol: true,
	}
	db, err := gorm.Open(postgres.New(pgConfig), ormConfig)
	if err != nil {
		return nil, err
	}

	s, err := db.DB()
	if err != nil {
		return nil, err
	}
	s.SetConnMaxLifetime(time.Minute * 5)
	s.SetMaxIdleConns(0)
	s.SetMaxOpenConns(20)

	return db, nil
}

func Close(db *gorm.DB) error {
	s, err := db.DB()
	if err != nil {
		return err
	}

	err = s.Close()
	if err != nil {
		return err
	}

	return nil
}
