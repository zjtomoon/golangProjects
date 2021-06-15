package main

import "runtime"

func main() {
	//获取当前的GOMAXPROCS
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	//设置GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)
	//获取当前的GOMAXPROCS
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
}
