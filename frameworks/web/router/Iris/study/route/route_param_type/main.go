package main

import "github.com/kataras/iris/v12"

// 路径参数类型
// Iris 拥有你见过的的最简单和强大路由处理。
//
// Iris 自己拥有用于路由路径语法解析和判定的解释器(就像一门编程语言)。
//
// 它是快速的。它计算它的需求，如果没有特殊的正则需要，它仅仅会使用低级的路径语法来注册路由，
// 除此之外，它预编译正则，然后加入到必需的中间件中。这就意味着与其他的路由器或者 web 框架相比，你的性能成本为零

func main() {
	app := iris.Default()

	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	app.Get("/users/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	app.Run(iris.Addr(":3000"))
}
