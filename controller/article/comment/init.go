package comment

import (
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/config"
)

func Init(g *echo.Group, cnf *config.Config) {
	cg := g.Group("/:article_id/comment")
	cg.GET("", GetList)
	cg.POST("", Add, auth.JWTWithConfig(cnf))
	cg.DELETE("/:id", Delete, auth.JWTWithConfig(cnf))
}
