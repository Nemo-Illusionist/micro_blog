package base

import (
	"github.com/labstack/echo"
	"micro_blog/core"
	"strconv"
)

type PageResponse struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Data     interface{} `json:"data"`
}

func NewPageResponse(data interface{}, page, pageSize int) PageResponse {
	pageResponse := PageResponse{
		Page:     page,
		PageSize: pageSize,
		Data:     data,
	}
	return pageResponse
}

func GetPaging(c echo.Context, ac *core.AppContext) (int, int) {

	pageSizeString := c.QueryParam("page_size")
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil || pageSize < 1 {
		if err != nil && pageSize > ac.Config.PageSizeMax {
			pageSize = ac.Config.PageSizeMax
		} else {
			pageSize = ac.Config.PageSize
		}
	}

	pageString := c.QueryParam("page")
	page, err := strconv.Atoi(pageString)
	if err != nil || page < 1 {
		page = 1
	}

	return page, (page - 1) * pageSize
}
