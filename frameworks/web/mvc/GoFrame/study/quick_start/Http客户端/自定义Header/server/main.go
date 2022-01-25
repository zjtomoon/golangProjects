package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writef(
			"Span-Id:%s,Trace-Id:%s",
			r.Header.Get("Span-Id"),
			r.Header.Get("Trace-Id"))
	})
	s.SetPort(8199)
	s.Run()
}
