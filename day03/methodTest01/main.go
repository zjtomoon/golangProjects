package main

import (
	"fmt"
)

type SliceInt []int

// func (s SliceInt) Sum int {
// 	sum := 0
// 	for _, i := range s {
// 		sum += i
// 	}
// 	return sum
// }
// //等价方法
// func SliceInt_Sum(s SliceInt) int {
// 	sum := 0
// 	for _, i := range s {
// 		sum += i
// 	}
// 	return sum
// }

type Map map[string]string

func (m Map) Print() {
	//底层类型支持的range运算，新类型可用
	for _, key := range m {
		fmt.Println(key)
	}
}

type MyInt int

func main() {
	var a MyInt = 10
	var b MyInt = 10
	//int 类型支持的加减乘除运算，新类型同样可用
	c := a + b
	d := a * b

	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
}
