package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"micro_blog/config"
	"micro_blog/dal/models"
	"time"
)

var (
	_authScheme = middleware.DefaultJWTConfig.AuthScheme
	_prefix     = _authScheme + " "
)

type jwtCustomClaims struct {
	UserID  uint64 `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenToken(user *models.User, cnf *config.Config) (string, error) {
	claims := &jwtCustomClaims{
		user.ID,
		user.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * cnf.TokenExpHour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(cnf.TokenSecretKey))
	return _prefix + t, err
}

func GetUserInfo(c echo.Context) (uint64, bool) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	userId := claims.UserID
	isAdmin := claims.IsAdmin
	return userId, isAdmin
}

func JWTWithConfig(cnf *config.Config) echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(cnf.TokenSecretKey),
		AuthScheme: _authScheme,
	}
	return middleware.JWTWithConfig(jwtConfig)
}
