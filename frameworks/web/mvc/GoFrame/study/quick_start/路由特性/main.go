package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/{class}--{course}/:name/*act", func(r *ghttp.Request) {
		r.Response.Writef(
			"%v %v %v %v",
			r.Get("class"),
			r.Get("course"),
			r.Get("name"),
			r.Get("act"),
		)
	})

	s.SetPort(8199)
	s.Run()
}
