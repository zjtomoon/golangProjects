package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	//redirect重定向中间件
	//https重定向
	e.Pre(middleware.HTTPSRedirect())

	//https www重定向
	e.Pre(middleware.HTTPSWWWRedirect())

	//https nonwww重定向
	e.Pre(middleware.HTTPSNonWWWRedirect())

	//www重定向
	e.Pre(middleware.WWWRedirect())

	//nonwww重定向
	e.Pre(middleware.NonWWWRedirect())

	e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
		Code: http.StatusTemporaryRedirect,
	}))
}
