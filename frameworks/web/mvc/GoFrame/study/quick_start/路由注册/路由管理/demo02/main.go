package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server().Domain("localhost")
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello World!")
	})
	g.Server().Run()
}
/*
	我们再次使用 http://127.0.0.1/ 进行访问，发现WebServer返回404，为什么呢？
	因为该程序中的回调函数只注册到了localhost域名中，因此只能通过 http://localhost/ 访问，其他域名自然无法访问。
 */