package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//rewrite 重写中间件 根据提供的规则重写URL路径，适用于向后兼容与创建更清晰、更具描述性的链接
	e.Pre(middleware.Rewrite(map[string]string{
		"/old":              "/new",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}))

	e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{}))
}
