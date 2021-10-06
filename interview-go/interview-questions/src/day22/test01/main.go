package main

import "fmt"

func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}

/*
	两个地方有语法问题。golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。
 */