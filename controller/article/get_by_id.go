package article

import (
	"github.com/labstack/echo"
	"micro_blog/controller/base"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
	"strconv"
)

func GetById(c echo.Context) error {
	ac := c.(*core.AppContext)

	articleID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	article := models.Article{ID: articleID}
	response := ac.Db.Model(&models.Article{}).Joins("User").
		Where(&models.Article{DeletedAt: nil}).
		First(&article)
	if response.Error != nil {
		return c.JSON(errors.NotFoundBoom())
	}

	var vm = Dto{
		ID:        article.ID,
		Title:     article.Title,
		Body:      article.Body,
		CreatedAt: article.CreatedAt,
		User: base.User{
			ID:   article.UserID,
			Name: article.User.Name,
		},
	}
	return c.JSON(http.StatusOK, vm)
}
