package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()

	//路由可以按照任意顺序定义
	e.GET("/hello", hello)
	e.GET("/users/id:", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})
	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})
	e.GET("/users/1/files", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/1/files/*")
	})
	//组路由
	//g := e.Group("/admin")

	//路由命名
	routeInfo := e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})
	routeInfo.Name = "user"
	//或者这样写
	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	}).Name = "newuser"

	//构造URI
	//业务处理
	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
	e.GET("/users/:id", h).Name = "foobar"

	//路由类别处理
	e.POST("/users", createUser)
	e.GET("/users", findUser)
	e.PUT("/users", updateUser)
	e.DELETE("/users", deleteUser)

	//用下面的代码将所有的路由信息输出到json文件
	data, _ := json.MarshalIndent(e.Routes(), "", " ")
	//if err != nil {
	//	return err
	//}
	ioutil.WriteFile("routes.json", data, 0644)

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

//路由列表
func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "createUser")
}

func findUser(c echo.Context) error {
	return c.String(http.StatusOK, "findUser")
}
func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "updateUser")
}
func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "deleteUser")
}
