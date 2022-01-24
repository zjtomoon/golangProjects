package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func MiddlewareAuth(r *ghttp.Request) {
	token := r.Get("token")
	if token == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器故障")
	}
}

func main() {
	s := g.Server()
	s.Use(MiddlewareCORS)
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareAuth, MiddlewareErrorHandler)
		group.ALL("/user/list", func(r *ghttp.Request) {
			panic("db error : sql is xxxxxxxx")
		})
	})
	s.SetPort(8199)
	s.Run()
}
