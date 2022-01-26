package main

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
)

func main() {
	data, err := gtcp.SendRecv("www.baidu.com:80", []byte("HEAD / HTTP/1.1\n\n"), -1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
