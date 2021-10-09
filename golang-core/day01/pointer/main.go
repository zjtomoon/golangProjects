package main

import (
	"fmt"
)

func main() {
	var a = 11
	p := &a
	fmt.Println(a, *p) //a和*p都为11

	type User struct {
		name string
		age  int
	}
	andes := User{
		name: "andes",
		age:  18,
	}
	q := &andes
	fmt.Println(q.name)
	/**
	go不支持指针的运算
	go 由于支持垃圾回收，如果支持指针运算，则会给垃圾回收的实现带来不便 禁止p++
	*/
	fmt.Println(sum(3, 5))
}
func sum(a, b int) *int {
	sum := a + b
	return &sum
}
