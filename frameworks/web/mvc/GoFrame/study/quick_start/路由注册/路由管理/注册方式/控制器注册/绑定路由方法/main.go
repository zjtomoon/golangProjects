package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) Index(r *ghttp.Request) {
	r.Response.Write("index")
}

func (c *Controller) Show(r *ghttp.Request) {
	r.Response.Write("show")
}

func main() {
	s := g.Server()
	c := new(Controller)
	//我们可以通过BindObjectMethod方法绑定指定的路由到指定的方法执行（方法名称参数区分大小写）。
	s.BindObjectMethod("/show", c, "Show")
	s.SetPort(3000)
	s.Run()
}
