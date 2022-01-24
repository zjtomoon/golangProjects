package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write(r.Get("map"))
	})
	s.SetPort(8199)
	s.Run()
}

// http://127.0.0.1:8199/?map[id]=1&map[name]=john
//http://127.0.0.1:8199/?map[user1][id]=1&map[user1][name]=john&map[user2][id]=2&map[user2][name]=smith
