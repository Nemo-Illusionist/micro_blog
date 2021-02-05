package article

import (
	"github.com/labstack/echo"
	"micro_blog/auth"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
	"strconv"
	"time"
)

func Delete(c echo.Context) error {
	_, isAdmin := auth.GetUserInfo(c)
	if !isAdmin {
		return c.JSON(errors.ForbiddenBoom())
	}

	articleID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	utc := time.Now().UTC()
	tx := c.(*core.AppContext).Db.Where(models.Article{ID: articleID}).Updates(models.Article{DeletedAt: &utc})
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.NoContent(http.StatusNoContent)
}
