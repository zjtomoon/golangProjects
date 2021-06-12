//函数签名
package main

import "fmt"


func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a-b

}
type Op func(int,int) int //定义一个函数类型，输入的是两个int类型，返回值是一个int类型

func do(f Op, a, b int) int {
	return f(a,b)
}

func main(){
	fmt.Printf("%T\n",add)
	a := do(add,1,2)
	fmt.Println(a)
	s := do(sub,1,2)
	fmt.Println(s)
}