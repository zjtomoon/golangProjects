package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//密钥认证中间件

	//e.Use(middleware.KeyAuth(func(key string) bool {
	//	return key == "valid-key"
	//}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "query:api-key",
	}))
}
