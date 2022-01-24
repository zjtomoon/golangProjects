package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writef("name: %v,pass: %v", r.Get("name"), r.Get("pass"))
	})
	s.SetPort(8199)
	s.Run()
}

//curl "http://127.0.0.1:8199/?name=john&pass=123"
//curl -d "name=john&pass=123" "http://127.0.0.1:8199/"
//curl -d '{"name":"john","pass":"123"}' "http://127.0.0.1:8199/"
//curl -d '<?xml version="1.0" encoding="UTF-8"?><doc><name>john</name><pass>123</pass></doc>' "http://127.0.0.1:8199/"
// curl -d '<doc><name>john</name><pass>123</pass></doc>' "http://127.0.0.1:8199/"
