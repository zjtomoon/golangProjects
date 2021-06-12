package main

import (
	"fmt"
)

func main() {
	s := "hello,世界~！"
	var a []byte
	var b string
	b = string(a)
	var c []rune
	c = []rune(s)
	fmt.Printf("%T\n", a) //[]uint8 byte是int8的別名
	fmt.Printf("%T\n", b) // string
	fmt.Printf("%T\n", c) // []int32 rune是int32的別名
}
