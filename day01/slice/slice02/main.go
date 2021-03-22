package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int, 2, 4)
	c := a[0:3]

	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(b, 1)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(b, c...)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	d := make([]int, 2, 3)
	copy(d, c) //copy只会复制d和C中长度最小的
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	str := "hello 世界！"
	e := []byte(str) //将字符串转化为[]byte类型切片
	f := []rune(str) //将字符类型转换为[]rune类型切片
	fmt.Println(e)
	fmt.Println(f)
}
