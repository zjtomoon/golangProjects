package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//使用日志中间件
	e.Use(middleware.Logger())

	//输出样例
	//{"time":"2017-01-12T08:58:07.372015644-08:00","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/","status":200, "latency":14743,"latency_human":"14.743µs","bytes_in":0,"bytes_out":2}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method},uri=${uri},status=${status}\n",
	}))

	//输出样例
	//method=GET, uri=/, status=200
}
