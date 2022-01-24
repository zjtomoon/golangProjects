package main

import (
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	total *gtype.Int
}

func (c *Controller) Total(r *ghttp.Request) {
	r.Response.Write("Total:", c.total.Add(1))
}

func main() {
	s := g.Server()
	c := &Controller{
		total: gtype.NewInt(),
	}

	s.BindHandler("/total", c.Total)
	s.SetPort(3000)
	s.Run()
}
