package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/url"
)

func main() {
	e := echo.New()

	url1, err := url.Parse("http://localhost:8081")
	if err != nil {
		e.Logger.Fatal(err)
	}

	url2, err := url.Parse("http://localhost:8082")
	if err != nil {
		e.Logger.Fatal(err)
	}
	//使用proxy代理中间件
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
	})))

	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{}))
}
