package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

// 构造函数
func (c *Controller) Init(r *ghttp.Request) {
	r.Response.Writeln("Init")
}

//析构函数
func (c *Controller) Shut(r *ghttp.Request) {
	r.Response.Writeln("Shut")
}

func (c *Controller) Hello(r *ghttp.Request) {
	r.Response.Writeln("Hello")
}

func main() {
	s := g.Server()
	c := new(Controller)
	s.BindObject("/object", c)
	s.SetPort(3000)
	s.Run()
}
