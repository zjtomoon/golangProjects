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

func main() {
	s := g.Server()
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.ALL("/login", func(r *ghttp.Request) {
			r.Response.Writeln("login")
		})
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareAuth)
			group.ALL("/dashboard", func(r *ghttp.Request) {
				r.Response.Writeln("dashboard")
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}
