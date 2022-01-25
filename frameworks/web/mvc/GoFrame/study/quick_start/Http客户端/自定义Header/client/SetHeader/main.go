package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func main() {
	c := g.Client()
	c.SetHeader("Span-Id", "0.0.1")
	c.SetHeader("Trace-Id", "NBC56410N97LJ016FQA")
	if r, e := c.Get("http://127.0.0.1:8199/"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}
