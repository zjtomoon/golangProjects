package main

import (
	"github.com/labstack/echo"
	"golangProjects/frameworks/Echo/study/inference/Context"
	"net/http"
)

func main() {
	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))

	//创建一个中间件来扩展默认的context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context.CustomContext{c}
			return h(cc)
		}
	})

	//在处理器中使用
	e.GET("/", func(c echo.Context) error {
		cc := c.(*Context.CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200, "OK")
	})

	//错误处理程序
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
	})

}
