package main

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func main() {
	path := "./main.go"
	r, e := ghttp.Post("http://127.0.0.1:8199/upload", "upload-file=@file:"+path)
	if e != nil {
		glog.Error(e)
	} else {
		fmt.Println(string(r.ReadAll()))
		r.Close()
	}
}
