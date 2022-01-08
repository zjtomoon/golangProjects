package router

import (
	"github.com/henrylee2cn/faygo"
	"golangProjects/frameworks/faygo/hello/handler"
	"golangProjects/frameworks/faygo/hello/middleware"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/", handler.Index),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),
	)
}
