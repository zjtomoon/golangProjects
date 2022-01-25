package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}

func main() {
	s := g.Server()
	s.Use(MiddlewareErrorHandler)
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.ALL("/user/list", func(r *ghttp.Request) {
			panic("db error: sql is xxxxxx")
		})
	})
	s.SetPort(8199)
	s.Run()
}
