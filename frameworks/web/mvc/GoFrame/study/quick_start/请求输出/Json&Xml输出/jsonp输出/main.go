package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//需要注意使用JSONP协议时必须通过Query方式提供callback参数。
func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/jsonp", func(r *ghttp.Request) {
			r.Response.WriteJsonP(g.Map{
				"id":   1,
				"name": "john",
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}

// curl -i "http://127.0.0.1:8199/jsonp?callback=MyCallback"
