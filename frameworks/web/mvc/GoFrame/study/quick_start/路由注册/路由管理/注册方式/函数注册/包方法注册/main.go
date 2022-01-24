package main

import (
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	total = gtype.NewInt()
)

func Total(r *ghttp.Request) {
	r.Response.Write("total:", total.Add(1))
}

func main() {
	s := g.Server()
	s.BindHandler("/total", Total)
	s.SetPort(3000)
	s.Run()
}
