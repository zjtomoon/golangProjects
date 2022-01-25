package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/json", func(r *ghttp.Request) {
			r.Response.WriteJson(g.Map{
				"id":   1,
				"name": "john",
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}

//curl -i http://127.0.0.1:8199/json
