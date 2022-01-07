package main

import "github.com/kataras/iris/v12"

func main() {
	//文件服务
	app := iris.New()

	app.HandleDir("/static", "./assets")

	//app.HandleDir("/static", "./assets", iris.DirOptions {
	//	Asset: Asset,
	//	AssetInfo: AssetInfo,
	//	AssetNames: AssetNames,
	//	Gzip: false,
	//})

	app.Run(iris.Addr(":3000"))
}
