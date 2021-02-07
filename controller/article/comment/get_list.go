package comment

import (
	"github.com/labstack/echo"
	"micro_blog/controller/base"
	"micro_blog/core"
	"micro_blog/core/errors"
	"micro_blog/dal/models"
	"net/http"
	"strconv"
)

func GetList(c echo.Context) error {
	ac := c.(*core.AppContext)

	articleID, err := strconv.ParseUint(c.Param("article_id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	page, offset := base.GetPaging(c, ac)

	var comments []models.Comment
	response := ac.Db.Model(&models.Comment{}).Joins("User").
		Where(&models.Comment{ArticleID: articleID}).
		Order("created_at desc").
		Offset(offset).Limit(page).
		Find(&comments)
	if response.Error != nil {
		return c.JSON(errors.InternalServerErrorBoom(response.Error))
	}

	list := make([]DtoElement, len(comments))

	for index, comment := range comments {
		list[index] = DtoElement{
			ID:        comment.ID,
			CommentID: comment.CommentID,
			CreatedAt: comment.CreatedAt,
			User: base.User{
				ID:   comment.ID,
				Name: comment.User.Name,
			},
		}
	}
	pageResponse := base.NewPageResponse(list, page, len(list))

	return c.JSON(http.StatusOK, pageResponse)
}
