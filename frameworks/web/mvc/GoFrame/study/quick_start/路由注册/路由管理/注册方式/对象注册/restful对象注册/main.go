package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	r.Response.Write("GET")
}

// RESTFul - POST
func (c *Controller) Post(r *ghttp.Request) {
	r.Response.Write("POST")
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Write("Delete")
}

// 该方法无法映射，将会无法访问到
func (c *Controller) Hello(r *ghttp.Request) {
	r.Response.Write("Hello")
}

func main() {
	s := g.Server()
	c := new(Controller)
	s.BindObject("/object", c)
	s.SetPort(3000)
	s.Run()
}
