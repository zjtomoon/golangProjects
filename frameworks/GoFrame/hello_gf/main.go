package main

import (
	_ "hello_gf/boot"
	_ "hello_gf/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
