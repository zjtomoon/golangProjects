package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)

	app.UseGlobal(before)
	app.UseGlobal(after)

	//你也可以使用 ExecutionRules 强制处理器在没有 ctx.Next() 的情况下完成执行。你可以这样做：

	//app.SetExecutionRules(iris.ExecutionRules{
	//	Done:iris.ExecutionOptions{Force:true},
	//})

	app.Run(iris.Addr(":3000"))
}

func before(ctx iris.Context) {

}

func after(ctx iris.Context) {

}

func indexHandler(ctx iris.Context) {
	ctx.HTML("<h1>Index</h1>")
}

func contactHandler(ctx iris.Context) {
	ctx.HTML("<h1>Contact</h1>")

	ctx.Next()
}
