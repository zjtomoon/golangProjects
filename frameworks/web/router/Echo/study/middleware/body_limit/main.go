package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//请求体限制中间件
	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{}))
}
