package main

import "fmt"

//函数声明语句，用于Go代码调用汇编代码
//func add(int,int) int

//func (int,int) int

//有名函数定义，函数名是add
func add(a,b int) int {
	return a + b
}
type ADD func(int,int) int

var g ADD = add
func main() {
	f := func(a,b int) int {
		return a + b
	}
	g(1,2)
	f(1,2)

	//f和add的函数签名相同
	fmt.Printf("%T\n",f) 	//func(int, int) int
	fmt.Printf("%T\n",add)  //func(int, int) int
}
