package ErrorHandler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

//自定义HTTP错误处理程序
func customHttpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
