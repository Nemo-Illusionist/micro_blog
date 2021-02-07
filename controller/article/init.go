package article

import (
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/config"
	"micro_blog/controller/article/comment"
)

func Init(g *echo.Group, cnf *config.Config) {
	ug := g.Group("/article")
	ug.GET("", GetList)
	ug.GET("/:id", GetById)

	ug.POST("", Add, auth.JWTWithConfig(cnf))
	ug.PUT("/:id", Update, auth.JWTWithConfig(cnf))
	ug.PATCH("/:id", HalfUpdate, auth.JWTWithConfig(cnf))
	ug.DELETE("/:id", Delete, auth.JWTWithConfig(cnf))

	comment.Init(ug, cnf)
}
