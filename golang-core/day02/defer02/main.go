package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	println("func body")
	os.Exit(1)
}

/*
	defer的好处是可以在一定程度上避免资源泄露。。特别是有很多return语句
	defer语句的位置不当,有可能导致panic,一般defer语句放在错误检查语句之后
	defer也有明显的副作用:defer会推迟资源的释放,defer尽量不要放到循环语句里面
*/
