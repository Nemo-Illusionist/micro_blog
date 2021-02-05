package auth

import (
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/config"
)

func Init(g *echo.Group, cfg *config.Config) {
	ug := g.Group("/auth")
	ug.POST("/login", Login)
	ug.POST("/registration", Registration)
	ug.POST("/change_password", ChangePassword, auth.JWTWithConfig(cfg))
}
