package article

import (
	"encoding/json"
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/controller/base"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
)

func Add(c echo.Context) error {
	userId, isAdmin := auth.GetUserInfo(c)
	if !isAdmin {
		return c.JSON(errors.ForbiddenBoom())
	}

	body := c.Request().Body
	dto := Request{}
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	ac := c.(*core.AppContext)
	article := &models.Article{
		UserID:    userId,
		Title:     dto.Title,
		ShortBody: dto.ShortBody,
		Body:      dto.Body,
	}

	tx := ac.Db.Create(&article)
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.JSON(http.StatusCreated, base.IdResponse{ID: article.ID})
}
