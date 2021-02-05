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

func ChangePassword(c echo.Context) error {
	body := c.Request().Body
	dto := PasswordRequest{}
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	userId, _ := auth.GetUserInfo(c)

	ac := c.(*core.AppContext)
	user := models.User{}
	user.SetPasswordHash(dto.Password, ac.Config.PasswordSecretKey)
	tx := ac.Db.Where(&models.User{ID: userId}).Updates(user)
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.NoContent(http.StatusNoContent)
}
