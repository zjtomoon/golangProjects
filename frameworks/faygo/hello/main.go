package main

import (
	"github.com/henrylee2cn/faygo"
	"golangProjects/frameworks/faygo/hello/router"
)

func main() {
	router.Route(faygo.New("hello"))
	faygo.Run()
}
