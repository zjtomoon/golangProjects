package main

import "github.com/gogf/gf/frame/g"

func main() {
	s := g.Server()
	s.SetServerRoot("/Users/john/Temp")
	s.SetRewrite("/test.html", "/test1.html")
	s.SetRewriteMap(g.MapStrStr{
		"/my-test1": "/test1.html",
		"/my-test2": "/test2.html",
	})
	s.SetPort(8199)
	s.Run()
}
