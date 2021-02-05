package main

import (
	"context"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"log"
	"micro_blog/config"
	"micro_blog/controller"
	"micro_blog/core"
	"micro_blog/dal"
	"micro_blog/migration"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logErr(err)
		return
	}

	db, err := dal.OpenWithConfig(cfg)
	if err != nil {
		logErr(err)
		return
	}

	err = migration.Migrate(db, migration.GetMigrationRange(), migration.Up)
	if err != nil {
		logErr(err)
		return
	}
	cont := &core.AppContext{
		Db:     db,
		Config: cfg,
	}

	router := core.CreateRouter(cont)
	controller.Init(router, cfg)

	go Start(router, cfg.Address)

	GracefulShutdown(db, router)
}

func Start(echo *echo.Echo, address string) {
	err := echo.Start(address)
	if err != nil {
		echo.Logger.Info("shutting down the server")
	}
}

func GracefulShutdown(db *gorm.DB, echo *echo.Echo) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := dal.Close(db)
	if err != nil {
		logErr(err)
	}

	err = core.CloseRouter(ctx, echo)
	if err != nil {
		logErr(err)
	}
}

func logErr(err error) {
	log.Fatalf("%+v\n", err)
}
