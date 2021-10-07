package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1,2,3}

	for i,v := range m {
		go func() {
			fmt.Println(i,v)
		}()
	}

	time.Sleep(3 * time.Second)
}

/*
	for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。

	各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，而不是各个goroutine启动时的i, v值。

	可以理解为闭包引用，使用的是上下文环境的值。
 */