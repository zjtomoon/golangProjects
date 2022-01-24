package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		if r.GetInt("type") == 1 {
			r.Response.Writeln("john")
			//加入退出方法 表示仅退出当前执行的逻辑方法，不退出后续的请求流程，可用于替代return。
			r.Exit()
		}
		r.Response.Writeln("smith")
	})
	s.SetPort(8199)
	s.Run()
}

//Exit: 仅退出当前执行的逻辑方法，不退出后续的请求流程，可用于替代return。
//ExitAll: 强行中断当前执行流程，当前执行方法的后续逻辑以及后续所有的逻辑方法将不再执行，常用于权限控制。
//ExitHook: 当路由匹配到多个HOOK方法时，默认是按照路由匹配优先级顺序执行HOOK方法。
// 当在HOOK方法中调用ExitHook方法后，后续的HOOK方法将不会被继续执行，作用类似HOOK方法覆盖。
//这三个退出函数仅在服务函数和HOOK事件回调函数中有效，无法控制中间件的执行流程。
//由于ExitAll和ExitHook方法在应用层比较少用，因此这里仅介绍Exit方法的使用。
