package main

import "github.com/gogf/gf/frame/g"

//静态服务
func main() {
	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot("/home/www/")
	s.Run()
}
