package main

import (
	_ "hello_beegoV2/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

