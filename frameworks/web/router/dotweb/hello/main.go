package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func main() {
	app := dotweb.New()
	app.SetLogPath("./logs/")

	app.HttpServer.GET("/index", func(ctx dotweb.Context) error {
		return ctx.WriteString("Welcome to my first web")
	})

	fmt.Println("dotweb.StartServer begin")
	err := app.StartServer(3000)
	fmt.Println("dotweb.StartServer error =>", err)
}
