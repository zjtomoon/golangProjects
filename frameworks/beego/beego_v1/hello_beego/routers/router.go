package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hello_beego/controllers"
)

func init() {

	//基础路由
	//基本get路由
	beego.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	//基本post路由
	beego.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	//注册一个可以响应任何http的路由
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})

	// restful controller路由
	//固定路由
	beego.Router("/", &controllers.MainController{})

	//初始化 namespace
	//ns :=
	//	beego.NewNamespace("/v1",
	//		beego.NSCond(func(ctx *context.Context) bool {
	//			if ctx.Input.Domain() == "api.beego.me" {
	//				return true
	//			}
	//			return false
	//		}),
	//		beego.NSBefore(auth),
	//		beego.NSGet("/notallowed", func(ctx *context.Context) {
	//			ctx.Output.Body([]byte("notAllowed"))
	//		}),
	//		beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
	//		beego.NSRouter("/changepassword", &UserController{}),
	//		beego.NSNamespace("/shop",
	//			beego.NSBefore(sentry),
	//			beego.NSGet("/:id", func(ctx *context.Context) {
	//				ctx.Output.Body([]byte("notAllowed"))
	//			}),
	//		),
	//		beego.NSNamespace("/cms",
	//			beego.NSInclude(
	//				&controllers.MainController{},
	//				&controllers.CMSController{},
	//				&controllers.BlockController{},
	//			),
	//		),
	//	)
	////注册 namespace
	//beego.AddNamespace(ns)

	//过滤器
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	//过滤器实现路由
	//var UrlManager = func(ctx *context.Context) {
	//	// 数据库读取全部的 url mapping 数据
	//	urlMapping := model.GetUrlMapping()
	//	for baseurl,rule:=range urlMapping {
	//		if baseurl == ctx.Request.RequestURI {
	//			ctx.Input.RunController = rule.controller
	//			ctx.Input.RunMethod = rule.method
	//			break
	//		}
	//	}
	//}
	//
	//beego.InsertFilter("/*",beego.BeforeRouter,UrlManager)

}
