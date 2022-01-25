package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func main() {
	c := g.Client()
	c.SetHeaderRaw(`
		Referer:https://localhost
		Span-Id: 0.0.1
		Trace-Id: NBC56410N97LJ016FQA
		User-Agent: MyTestClient
	`)

	if r, e := c.Get("http://127.0.0.1:8199/"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}
