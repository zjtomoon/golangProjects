package main

import "github.com/kataras/iris/v12"

// 反向查询
func main() {
	app := iris.New()
	h := func(ctx iris.Context) {
		ctx.HTML("<b>Hi</b>")
	}

	//	路由命名非常简单，我们只需要调用返回的 *Route 的 Name 字段来定义名字。
	home := app.Get("/", h)
	home.Name = "home"

	app.Get("/about", h).Name = "about"
	app.Get("/page/{id}", h).Name = "page"

	app.Run(iris.Addr(":3000"))
}
