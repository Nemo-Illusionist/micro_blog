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

func Registration(c echo.Context) error {
	ac := c.(*core.AppContext)
	body := c.Request().Body
	dto := RegistrationRequest{}
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	user := models.User{Name: dto.Name, Login: dto.Login, IsAdmin: false}
	user.SetPasswordHash(dto.Password, ac.Config.PasswordSecretKey)

	tx := ac.Db.Create(&user)
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	token, err := auth.GenToken(&user, ac.Config)
	if err != nil {
		return c.JSON(errors.InternalServerErrorBoom(err))
	}

	return c.JSON(http.StatusOK, RegistrationResponse{UserId: user.ID, LoginResponse: LoginResponse{Token: token}})
}
