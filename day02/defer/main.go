package main

import "fmt"

//defer 后面必须是函数或方法的调用，不能是语句.否则会报expression in defer must be function call错误
func main() {
	//先进后出
	defer func() {
		fmt.Println("first") //no.3
	}()
	defer func() {
		fmt.Println("second") //no.2
	}()
	fmt.Println("function body") //no.1
	f()
}
func f() int {
	a := 0
	defer func(i int) {
		fmt.Println("defer i=", i)
	}(a)
	a++
	return a
}
