package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//static 静态中间件  Static 中间件可已被用于服务从根目录中提供的静态文件。
	e.Use(middleware.Static("/static"))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	}))
}
