package main

import (
	"fmt"
	"log"
	"net/http"
)

type greetingHandler struct {
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,World")
}

func (h greetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //写成ServerHTTP会报错
	fmt.Fprintf(w, "Hello,%s", h.Name)
}

func main() {
	mux := http.NewServeMux()
	//注册处理器函数
	mux.HandleFunc("/hello", helloHandler)

	//注册处理器
	mux.Handle("greeting/golang", greetingHandler{Name: "Golang"}) //?

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
