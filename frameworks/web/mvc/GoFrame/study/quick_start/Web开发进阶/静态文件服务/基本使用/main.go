package main

import "github.com/gogf/gf/frame/g"

func main() {
	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot("/Users/john/Temp")
	s.AddSearchPath("/Users/john/Documents")
	s.SetPort(8199)
	s.Run()
}
