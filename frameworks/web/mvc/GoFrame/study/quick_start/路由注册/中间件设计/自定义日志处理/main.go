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

func MiddlewareLog(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := ""
	if err := r.GetError(); err != nil {
		errStr = err.Error()
	}
	g.Log().Println(r.Response.Status, r.URL.Path, errStr)
}

func main() {
	s := g.Server()
	s.SetConfigWithMap(g.Map{
		"AccessLogEnabled": false,
		"ErrorLogEnabled":  false,
	})
	s.Use(MiddlewareLog, MiddlewareCORS)
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareAuth)
		group.ALL("/user/list", func(r *ghttp.Request) {
			panic("啊！我出错了！")
		})
	})

	s.SetPort(8199)
	s.Run()
}
