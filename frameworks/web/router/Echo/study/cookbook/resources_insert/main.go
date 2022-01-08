package main

import (
	rice "github.com/GeertJohan/go.rice" //使用go.rice
	"github.com/labstack/echo"
	"net/http"
)

//资源嵌入
func main() {
	e := echo.New()

	assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())

	e.GET("/", echo.WrapHandler(assetHandler))

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	e.Logger.Fatal(e.Start(":1323"))
}
