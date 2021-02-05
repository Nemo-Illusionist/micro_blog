package errors

import (
	"github.com/labstack/echo"
)

func HTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	switch v := err.(type) {
	case *echo.HTTPError:
		err := c.JSON(NewBoom(v.Code, "", v))
		if err != nil {
			c.Logger().Error("error handler: json encoding", err)
		}
	default:
		err := c.JSON(InternalServerErrorBoom(v))
		if err != nil {
			c.Logger().Error("error handler: json encoding", err)
		}
	}
}
