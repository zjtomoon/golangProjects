package main

import "fmt"

//函数声明语句，用于Go代码调用汇编代码
//func add(int,int) int

//func (int,int) int

//有名函数定义，函数名是add
func add(a,b int) int {
	return a + b
}
//新定义函数类型ADD
//ADD底层类型是函数字面值类型 func (int,int) int
type ADD func(int,int) int

//add和ADD的底层类型相同，并且add是字面量类型
//所以add可以直接赋值给ADD类型的变量g
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
