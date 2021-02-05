package core

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
	"micro_blog/config"
	"micro_blog/core/errors"
)

type AppContext struct {
	echo.Context
	Db     *gorm.DB
	Config *config.Config
}

func (cc *AppContext) ToMiddlewareFunc() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc.Context = c
			return h(cc)
		}
	}
}

func CreateRouter(context *AppContext) *echo.Echo {
	e := echo.New()
	e.Use(context.ToMiddlewareFunc())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("5M"))
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS, echo.HEAD},
	}))

	e.HTTPErrorHandler = errors.HTTPErrorHandler

	return e
}

func CloseRouter(ctx context.Context, e *echo.Echo) error {

	err := e.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
