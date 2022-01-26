package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gproc"
	"time"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Hello! ")
	})
	s.BindHandler("/pid", func(r *ghttp.Request) {
		r.Response.Writeln(gproc.Pid())
	})
	s.BindHandler("/sleep", func(r *ghttp.Request) {
		r.Response.Writeln(gproc.Pid())
		time.Sleep(10 * time.Second)
		r.Response.Writeln(gproc.Pid())
	})
	s.EnableAdmin()
	s.SetPort(8199)
	s.Run()
}

//访问 http://127.0.0.1:8199/pid 查看当前进程的pid
//访问 http://127.0.0.1:8199/sleep ，这个页面将会执行10秒，用于测试重启时该页面请求执行是否会断掉
//访问 http://127.0.0.1:8199/debug/admin ，这是s.EnableAdmin后默认注册的一个WebServer管理页面
