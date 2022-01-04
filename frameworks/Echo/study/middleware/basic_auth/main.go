package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	//使用自定义配置
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))
}

/*
	BasicAuthConfig struct {
  // Skipper 定义了一个跳过中间件的函数
  Skipper Skipper

  // Validator 是一个用来验证 BasicAuth 是否合法的函数
  // Validator 是必须的.
  Validator BasicAuthValidator

  // Realm 是一个用来定义 BasicAuth 的 Realm 属性的字符串
  // 默认值是 "Restricted"
  Realm string
}
*/
