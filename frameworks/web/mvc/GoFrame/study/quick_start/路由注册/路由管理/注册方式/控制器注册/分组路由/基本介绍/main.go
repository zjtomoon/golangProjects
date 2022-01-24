package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	//创建分组路由
	group := s.Group("/api")
	group.ALL("/all", func(r *ghttp.Request) {
		r.Response.Write("all")
	})

	group.GET("/get", func(r *ghttp.Request) {
		r.Response.Write("get")
	})

	group.POST("/post", func(r *ghttp.Request) {
		r.Response.Write("post")
	})

	s.SetPort(8199)
	s.Run()
}

//其中，/api/get仅允许GET方式访问，/api/post仅允许POST方式访问，/api/all允许所有的方式访问。
