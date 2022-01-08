package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", before, mainHandler, after)
	app.Run(iris.Addr(":3000"))
}

func before(ctx iris.Context) {
	shareInfomation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInfomation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	info := ctx.Values().GetString("info")
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("</br> Info: " + info)

	ctx.Next()
}
