package comment

import (
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
	"strconv"
)

func Delete(c echo.Context) error {
	userId, _ := auth.GetUserInfo(c)

	commentID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}
	articleID, err := strconv.ParseUint(c.Param("article_id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	where := models.Comment{ID: commentID, UserID: userId, ArticleID: articleID}
	tx := c.(*core.AppContext).Db.Where(where).Delete(models.Comment{})
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.NoContent(http.StatusNoContent)
}
