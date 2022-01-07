package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":3000"))
}

// 包装路由器
func newApp() *iris.Application {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})
	app.Get("/profile/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("username"))
	})

	app.HandleDir("/", "./public")

	myOtherHandler := func(ctx iris.Context) {
		ctx.Writef("inside a handler which is fired manually by our custom router wrapper")
	}

	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		if strings.HasPrefix(path, "/other") {
			ctx := app.ContextPool.Acquire(w, r)
			myOtherHandler(ctx)
			app.ContextPool.Release(ctx)
			return
		}
	})
	return app
}
