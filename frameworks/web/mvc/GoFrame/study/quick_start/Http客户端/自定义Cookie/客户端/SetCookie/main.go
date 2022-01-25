package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func main() {
	c := g.Client()
	c.SetCookie("name", "john")
	c.SetCookie("score", "100")
	if r, e := c.Get("http://127.0.0.1:8199/"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}
