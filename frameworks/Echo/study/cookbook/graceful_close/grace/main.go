package main

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"net/http"
)

// 平滑关闭
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Six sick bricks tick")
	})

	e.Server.Addr = ":1323"

	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
