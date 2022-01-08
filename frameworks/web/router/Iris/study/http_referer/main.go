package main

import "github.com/kataras/iris/v12"

//HTTP referer (本来是 referrer 的拼写错误) 是一个可选的 HTTP 头部字段， 用于标记链接到请求资源的网页的地址(即 URI 或者 IRI)。通过检查 referrer，新的网页可以知道请求的来源
//ris 使用 Shopify's goreferrer 包来实现 Context.GetReferrer() 方法。
//
//GetReferrer 方法提取和返回 Referer 头的信息，Referer

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		r := ctx.GetReferrer()
		switch r.Type {
		case iris.ReferrerSearch:
			ctx.Writef("Search %s: %s\n", r.Label, r.Query)
			ctx.Writef("Google: %s\n", r.GoogleType)
		case iris.ReferrerSocial:
			ctx.Writef("Social %s\n", r.Label)
		case iris.ReferrerIndirect:
			ctx.Writef("Indirect: %s\n", r.URL)
		}
	})

	app.Run(iris.Addr(":3000"))
}
