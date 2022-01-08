package main

import "github.com/kataras/iris/v12"

//处理HTTP错误
func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	app.Get("/", index)
	app.Run(iris.Addr(":3000"))
}

func notFound(ctx iris.Context) {
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong,try again")
}

func index(ctx iris.Context) {
	ctx.View("index.html")
}

//Iris 内建支持 HTTP APIs 的错误详情。
//
//Context.Problem 编写一个 JSON 或者 XML 问题响应，行为完全类似 Context.JSON，但是默认 ProblemOptions.JSON 的缩进是 " "，响应的 Content-type 为 application/problem+json。
//
//使用 options.RenderXML 和 XML 字段来改变他的行为，用 application/problem+xml 的文本类型替代。
func newProductProblem(productNamem, detail string) iris.Problem {
	return iris.NewProblem().
		Type("/product-error").
		Title("Product validation problem").
		Detail(detail).
		Status(iris.StatusBadRequest).
		Key("productName", productNamem)
}

func fireProblem(ctx iris.Context) {
	ctx.Problem(newProductProblem("product name", "product details"),
		iris.ProblemOptions{
			JSON: iris.JSON{
				Indent: " ",
			},
			RetryAfter: 300,
		})
}
