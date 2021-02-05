package auth

import (
	"encoding/json"
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
)

func Login(c echo.Context) error {
	ac := c.(*core.AppContext)
	body := c.Request().Body
	dto := LoginRequest{}
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	user := models.User{Login: dto.Login, DeletedAt: nil}

	tx := ac.Db.Where(&user).First(&user)
	if tx.Error != nil || !user.CheckPassword(dto.Password, ac.Config.PasswordSecretKey) {
		return c.JSON(errors.NotFoundBoom())
	}

	token, err := auth.GenToken(&user, ac.Config)
	if err != nil {
		return c.JSON(errors.InternalServerErrorBoom(err))
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}
