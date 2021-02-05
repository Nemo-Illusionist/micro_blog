package controller

import (
	"github.com/labstack/echo"
	"micro_blog/config"
	"micro_blog/controller/article"
	"micro_blog/controller/auth"
)

func Init(e *echo.Echo, cnf *config.Config) {
	g := e.Group("/api")
	article.Init(g, cnf)
	auth.Init(g, cnf)
}
