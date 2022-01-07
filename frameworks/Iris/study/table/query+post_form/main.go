package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		app.Logger().Infof("id : #s; page : %s; name : %s; message : %s",
			id, page, name, message)
	})

	app.Run(iris.Addr(":3000"))
}
