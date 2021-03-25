package main

import (
	"fmt"
)

//匿名函数被直接赋值函数变量
var sum = func(a, b int) int {
		return a + b
	}
func doinput(f func(int,int) int , a ,b int) int {
	return f(a,b)
}
//匿名函数作为返回值
func wrap(op string) func(int,int)  int {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "sub":
		return func(a, b int) int  {
			return a + b
		}
	default:
		return nil
	}
}
func main(){
	//匿名函数直接被调用
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()
	sum(1,2)
	//匿名函数作为实参
	doinput(func(x int, y int) int {
		return x + y
	},1,2)

	opFunc := wrap("add")
	res := opFunc(2,3)
	fmt.Printf("%d\n",res)
}