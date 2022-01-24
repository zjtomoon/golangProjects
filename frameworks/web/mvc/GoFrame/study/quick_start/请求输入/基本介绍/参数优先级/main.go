package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/input", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("amount"))
	})
	s.BindHandler("/query", func(r *ghttp.Request) {
		r.Response.Writeln(r.GetQuery("amount"))
	})
	s.SetPort(8199)
	s.Run()
}

/*
	$ curl -d "amount=1" -X POST "http://127.0.0.1:8199/input?amount=100"
	1
	$ curl -d "amount=1" -X POST "http://127.0.0.1:8199/query?amount=100"
	100
*/

//可以看到，当我们访问/input路由时，该路由方法中采用了Get*方法获取amount参数，按照同名优先级的规则返回了1，
// 即body中传递的参数。而当我们通过/query路由访问时，
// 该路由方法中使用了GetQuery*方法获取amount参数，因此获取到的是query string参数中的amount值，返回了100
