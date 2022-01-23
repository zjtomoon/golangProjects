package main

import "github.com/gogf/gf/frame/g"

//多实例支持
func main() {
	s1 := g.Server("s1")
	s1.SetPort(8080)
	s1.SetIndexFolder(true)
	s1.SetServerRoot("/home/www/static1")
	s1.Start()

	s2 := g.Server("s2")
	s2.SetPort(8088)
	s2.SetIndexFolder(true)
	s2.SetServerRoot("/home/www/static2")
	s2.Start()

	g.Wait()
}
