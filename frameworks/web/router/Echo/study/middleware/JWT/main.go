package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// JWT (json web token)认证中间件
	e.Use(middleware.JWT([]byte("secret")))

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "quert:token",
	}))
}
