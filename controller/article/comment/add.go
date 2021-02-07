package comment

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

func Add(c echo.Context) error {
	userId, _ := auth.GetUserInfo(c)
	articleId, err := strconv.ParseUint(c.Param("article_id"), 10, 0)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	body := c.Request().Body
	dto := Request{}
	err = json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	ac := c.(*core.AppContext)
	comment := &models.Comment{
		UserID:    userId,
		ArticleID: articleId,
		Body:      dto.Body,
		CommentID: dto.CommentID,
	}

	tx := ac.Db.Create(&comment)
	if tx.Error != nil {
		return c.JSON(errors.BadRequestBoom(err))
	}

	return c.JSON(http.StatusCreated, base.IdResponse{ID: comment.ID})
}
