package main

import "github.com/kataras/iris/v12"

//子域名

//Iris 具有已知最简单的注册子域名到单个应用程序的方式。当然你也可以在生成环境中使用 nginx 或者 caddy。

// 子域名被分为两类：静态 和 动态/通配。
//
//静态：你所知的子域名，例如：analytics.mydomain.com。
//通配：翻译不通(when you don’t know the subdomain but you know that it’s before a particular subdomain or root domain, i.e : )，
// 即：user_created.mydomain.com，otheruser.mydomain.com，就像 username.github.io 这样。

func main() {
	app := iris.New()

	admin := app.Subdomain("admin")

	admin.Get("/", func(ctx iris.Context) {
		ctx.Writef("INDEX FROM admin.mydomain.com")
	})

	admin.Get("/hey", func(ctx iris.Context) {
		ctx.Writef("HEY FROM admin.mydomain.com/hey")
	})

	//为了证明子域名像其它正则 Party 一样工作，你也可以用下面这种另类的方法注册一个子域名：

	//adminSubdomain:= app.Party("admin.")
	//// or
	//adminAnalayticsSubdomain := app.Party("admin.analytics.")
	//// or for a dynamic one:
	//anySubdomain := app.Party("*.")

	//上面的所有 htt(s)://mydomain.com/%anypath% 将会重定向到 https(s)://www.mydomain.com/%anypath%。
	www := app.Subdomain("www")
	app.SubdomainRedirect(app, www)

	app.Run(iris.Addr("mydomain.com:80"))
}

func info(ctx iris.Context) {
	method := ctx.Method()
	subdomain := ctx.Subdomain()
	path := ctx.Path()

	ctx.Writef("\nInfo\n\n")
	ctx.Writef("Method: %s\nSubdomain: %s\nPath: %s", method, subdomain, path)
}
