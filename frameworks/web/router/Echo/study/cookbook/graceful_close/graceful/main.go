package main

import (
	"github.com/labstack/echo"
	"github.com/tylerb/graceful"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sue sews rose on slow joe crows nose")
	})
	e.Server.Addr = ":1323"

	graceful.ListenAndServe(e.Server, 5*time.Second)
}
