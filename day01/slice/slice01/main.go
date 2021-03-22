package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type slice struct {
		array unsafe.Pointer
		len   int
		cat   int
	}
	var array = [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := array[0:4]
	s2 := array[:4]
	s3 := array[2:]
	fmt.Printf("%v\n",s1)
	fmt.Printf("%v\n",s2)
	fmt.Printf("%v\n",s3)
	//make创建的切片各元素被默认初始化为元素类型的零值
	a := make([]int, 10)

	//len = 10 , cap = 15
	b := make([]int, 10, 15)
	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	var c []int
	fmt.Printf("%v\n",c) // []

}
