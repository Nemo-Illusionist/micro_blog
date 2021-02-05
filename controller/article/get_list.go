package article

import (
	"github.com/labstack/echo"
	"micro_blog/controller/base"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
)

func GetList(c echo.Context) error {
	ac := c.(*core.AppContext)

	page, offset := base.GetPaging(c, ac)

	var articles []models.Article
	response := ac.Db.Model(&models.Article{}).Joins("User").
		Where(&models.Article{DeletedAt: nil}).
		Offset(offset).Limit(page).
		Find(&articles)
	if response.Error != nil {
		return c.JSON(errors.InternalServerErrorBoom(response.Error))
	}

	list := make([]DtoElement, len(articles))

	for index, article := range articles {
		list[index] = DtoElement{
			ID:        article.ID,
			Title:     article.Title,
			ShortBody: article.ShortBody,
			CreatedAt: article.CreatedAt,
			User: User{
				ID:   article.ID,
				Name: article.User.Name,
			},
		}
	}
	pageResponse := base.NewPageResponse(list, page, len(list))

	return c.JSON(http.StatusOK, pageResponse)
}
