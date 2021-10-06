package main

import "fmt"

type S struct {
	m string
}

func f() *S {
	return &S{"foo"}
}

func main() {
	//p := f()
	p := *f()
	fmt.Println(p.m)
}

/*
	f() 函数返回参数是指针类型，所以可以用 & 取结构体的指针；B 处，如果填 *f()，
	则 p 是 S 类型；如果填 f()，则 p 是 *S 类型，不过都可以使用 p.m 取得结构体的成员。
 */