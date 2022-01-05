package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//requesr id 请求id中间件 为请求生成唯一的id
	e.Use(middleware.RequestID())

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return customGenerator()
		},
	}))
}

func customGenerator() string {

}
