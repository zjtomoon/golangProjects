package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//使用请求体转储中间件
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

	}))

	//使用自定义配置
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))
}

//BodyDumpConfig struct {
//// Skipper 定义了一个跳过中间件的函数
//Skipper Skipper
//
//// Handler 接收请求和响应有效负载
//// Required.
//Handler BodyDumpHandler
//}
