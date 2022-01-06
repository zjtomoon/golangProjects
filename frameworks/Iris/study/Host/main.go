package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	//app.Run(iris.Addr(":3000"),iris.WithTunneling)

	//使用配置
	//config := iris.WithConfiguration(iris.Configuration{
	//	DisableStartupLog:true,
	//	EnableOptimizations:true,
	//	Charset:"UTF-8",
	//})
	//
	//app.Run(iris.Addr(":3030"),config)

	//从YAML加载
	//config := iris.WithConfiguration(iris.YAML("./iris.yml"))
	//app.Run(iris.Addr(":3000"),config)

	//从TOML加载
	//config := iris.WithConfiguration(iris.TOML("./iris.toml"))
	//app.Run(iris.Addr(":3000"),config)

	//使用函数式加载
	app.Run(iris.Addr(":3000"),
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithoutAutoFireStatusCode,
		iris.WithOptimizations,
		iris.WithTimeFormat("Mon, 01 Jan 15:04:05 GMT"))

	//自定义值
	app.Run(iris.Addr(":8080"),
		iris.WithOtherValue("ServerName", "my amazing iris server"),
		iris.WithOtherValue("ServerOwner", "admin@example.com"),
	)
	//通过app.ConfigurationReadOnly来访问以上字段
	//serverName := app.ConfigurationReadOnly().Other["MyServerName"]
	//serverOwner := app.ConfigurationReadOnly().Other["ServerOwner"]

	//app.Run(iris.Addr(":3000"),iris.WithConfiguration(iris.Configuration{
	//	Tunneling:iris.TunnelingConfiguration{
	//		AuthToken:"my-ngrok-auth-client-token",
	//		Bin:"$GOPATH/bin/ngrok",
	//		Region:"eu",
	//		WebInterface:"127.0.0.1:4040",
	//		Tunnels:[]iris.Tunnel{
	//			{
	//				Name:"MyApp",
	//				Addr:":8080",
	//			},
	//		},
	//	},
	//}))
}
