package main

import "github.com/kataras/iris/v12"

type testdata struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

//有时候一个服务器应用程序需要在相同的 URI 上提供资源不同表示形式。当然这也可以手动实现，检查 Accept 请求头，
// 设置推送的请求表单的文本。然而，当你的程序管理更多的资源和不同形式的时候，这将变得非常痛苦，
// 你需要检查 Accept-Charset，Accept-Encoding，设置一些服务器端优先级，正确处理错误等等。

func main() {
	app := iris.New()

	//通过 gzip 编码算法将资源渲染为 application/json 或者 text/html 或者 application/xml。
	//
	//当客户端的 Accept 头部包含上述之一时，
	//如果接收的为空就声明为 JSON(首先声明)
	//当客户端的 Accpet-Encoding 头包含 gzip 或者为空时
	app.Get("/resource", func(ctx iris.Context) {
		data := testdata{
			Name: "test name",
			Age:  26,
		}

		ctx.Negotiation().JSON().XML().EncodingGzip()

		_, err := ctx.Negotiate(data)
		if err != nil {
			ctx.Writef("%v", err)
		}

		// 或者在一个中间件中定义他们，然后在最后的处理器中以 nil 参数调用Negotiate。

		ctx.Negotiation().JSON(data).XML(data).Any("content for */*")
		ctx.Negotiate(nil)

		app.Get("/resource2", func(ctx iris.Context) {
			jsonAndXML := testdata{
				Name: "test name",
				Age:  26,
			}
			ctx.Negotiation().
				JSON(jsonAndXML).
				XML(jsonAndXML).
				HTML("<h1>Test Name</h1><h2>Age 26</h2>")

			ctx.Negotiate(nil)
		})
	})
}
