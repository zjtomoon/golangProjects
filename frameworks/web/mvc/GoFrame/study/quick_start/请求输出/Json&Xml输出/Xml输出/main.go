package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/xml", func(r *ghttp.Request) {
			r.Response.Write(`<?xml version="1.0" encoding="UTF-8"?>`)
			r.Response.WriteXml(g.Map{
				"id":   1,
				"name": "john",
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}

//curl -i http://127.0.0.1:8199/xml
