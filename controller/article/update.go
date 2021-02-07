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
	"strconv"
)

func Update(c echo.Context) error {
	_, isAdmin := auth.GetUserInfo(c)
	if !isAdmin {
		return c.JSON(errors.ForbiddenBoom())
	}

	body := c.Request().Body
	dto := Request{}
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	articleID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	ac := c.(*core.AppContext)
	article := &models.Article{Title: dto.Title, ShortBody: dto.ShortBody, Body: dto.Body}
	tx := ac.Db.Where(&models.Article{ID: articleID}).Updates(article)
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.JSON(http.StatusCreated, base.IdResponse{ID: articleID})
}
