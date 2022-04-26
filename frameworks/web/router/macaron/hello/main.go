package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
)

func main() {
	m := macaron.Classic()

	//m.Get("/", func() string{
	//	return "hello world !"
	//})
	//m.Run()
	m.Get("/", myHandler)
	log.Println("server is running...")
	log.Println(http.ListenAndServe(":8080", m))
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is :" + ctx.Req.RequestURI
}
