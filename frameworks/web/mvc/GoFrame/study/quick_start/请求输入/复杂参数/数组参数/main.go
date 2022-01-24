package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write(r.Get("array"))
	})
	s.SetPort(8199)
	s.Run()
}

// http://127.0.0.1:8199/?array[]=john&array[]=smith
