package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//多端口监听
func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("go frame!")
	})
	s.SetPort(8100, 8200, 8300)
	s.Run()
}
