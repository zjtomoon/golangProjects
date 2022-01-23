package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//绑定域名

func Hello1(r *ghttp.Request) {
	r.Response.Write("127.0.0.1: Hello1!")
}

func Hello2(r *ghttp.Request) {
	r.Response.Write("localhost: Hello2!")
}

func main() {
	s := g.Server()
	s.Domain("127.0.0.1").BindHandler("/", Hello1)
	s.Domain("localhost").BindHandler("/", Hello2)
	s.Run()
}
