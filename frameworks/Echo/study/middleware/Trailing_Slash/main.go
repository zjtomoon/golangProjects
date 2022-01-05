package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	//添加尾部斜杠中间件 Add trailing slash 中间件会在在请求的 URI 后加上反斜杠
	e.Pre(middleware.AddTrailingSlash())

	//去除尾部斜杠
	e.Pre(middleware.RemoveTrailingSlash())

	//自定义配置
	e.Pre(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}
