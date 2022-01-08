package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	// 创建echo实例
	e := echo.New()

	//使用中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 路由
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!\n")
	})

	// 开启服务
	e.Logger.Fatal(e.Start(":3000"))
}
