package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Order struct {
}

func (o *Order) List(r *ghttp.Request) {
	r.Response.Write("list")
}

func main() {
	s := g.Server()
	o := new(Order)
	s.BindObject("/{.struct}-{.method}", o)
	s.SetPort(3000)
	s.Run()
}
