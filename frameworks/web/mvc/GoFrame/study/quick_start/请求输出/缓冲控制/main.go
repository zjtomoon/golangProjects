package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

//Response输出采用了缓冲控制，输出的内容预先写入到一块缓冲区，等待服务方法执行完毕后才真正地输出到客户端。
// 该特性在提高执行效率同时为输出内容的控制提供了更高的灵活性。

//证明
//我们通过后置中间件统一对返回的数据做处理，如果服务方法产生了异常，
// 那么不能将敏感错误信息推送到客户端，而统一设置错误提示信息。

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}

func main() {
	s := g.Server()
	s.Group("api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareErrorHandler)
		group.ALL("/user/list", func(r *ghttp.Request) {
			//r.Response.Writeln("hello") // 正常输出hello
			panic("db err:sql is xxxxxx")
		})
	})
	s.SetPort(8199)
	s.Run()
}
