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

func HalfUpdate(c echo.Context) error {
	_, isAdmin := auth.GetUserInfo(c)
	if !isAdmin {
		return c.JSON(errors.ForbiddenBoom())
	}

	body := c.Request().Body
	dto := PatchRequest{}
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	articleID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	ac := c.(*core.AppContext)
	article := &models.Article{}
	if dto.Title != nil {
		article.Title = *dto.Title
	}

	if dto.Body != nil {
		article.Body = *dto.Body
	}

	if dto.ShortBody != nil {
		article.ShortBody = *dto.ShortBody
	}

	tx := ac.Db.Where(&models.Article{ID: articleID}).Updates(article)
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.JSON(http.StatusCreated, base.IdResponse{ID: articleID})
}
