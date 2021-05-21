package main

import (
	"fmt"
	"net/http")


func myhandler(w http.ResponseWriter,r *http.Request) {
	fmt.Println(r.RemoteAddr,"连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("Method:",r.Method)
	fmt.Println("url:",r.URL)
	fmt.Println("Header:",r.Header)
	fmt.Println("Body:",r.Body)

	//回复
	w.Write([]byte("www.51mh.com"))
}

func main() {
	http.HandleFunc("/go",myhandler)
	http.ListenAndServe("127.0.0.1:8000",nil)
}