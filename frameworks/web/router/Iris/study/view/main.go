package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	tmpl := iris.HTML("./views", ".html")

	tmpl.Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hi.html")
	})

	app.Run(iris.Addr(":3000"))
}
