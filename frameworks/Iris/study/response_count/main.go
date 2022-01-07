package main

import "github.com/kataras/iris/v12"

//在发送数据前调用 Context.Record
//Context.Recorder() 返回一个 ResponseRecorder。它的方法可以用来操作和找回响应。
//ResponseRecorder 类型包含了标准的 Iris ResponseWriter 方法和下列的方法。
//
//Body 返回目前为止 Writer 写入的响应体数据。不要使用它来编辑响应体。

//在一个全局拦截器中记录操作日志。
func main() {
	app := iris.New()
	// start record.
	app.Use(func(ctx iris.Context) {
		ctx.Record()
		ctx.Next()
	})

	app.Done(func(ctx iris.Context) {
		body := ctx.Recorder().Body()

		app.Logger().Infof("sent: %s", string(body))
	})

	app.Get("/save", func(ctx iris.Context) {
		ctx.WriteString("success")
		ctx.Next()
	})

	//或者为了消除你的主处理器中对 ctx.Next 的需求，改变 Iris 处理器的执行规则，你可以如下所示：
	app.SetExecutionRules(iris.ExecutionRules{
		Done: iris.ExecutionOptions{
			Force: true,
		},
	})

	app.Get("/save", func(ctx iris.Context) {
		ctx.WriteString("success")
	})
}
