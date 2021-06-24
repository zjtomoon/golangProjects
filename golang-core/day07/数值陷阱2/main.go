package main

import "fmt"

/*
	Go语言数组的一切传递都是值拷贝，体现在以下三个方面
	(1)数组间的直接赋值
	(2)数组作为函数参数
	(3)数组内嵌到struct中。
 */

func f(a [3]int)  {
	a[2] = 10
	fmt.Printf("%p,%v\n",&a,a)
}

func main() {
	a := [3]int{1,2,3}
	//直接赋值是值拷贝
	b := a

	//修改a元素值并不影响b
	a[2] = 4

	fmt.Printf("%p,%v\n",&a,a)
	fmt.Printf("%p,%v\n",&b,b)

	//数组作为函数参数仍然是值拷贝
	f(a)

	c := struct {
		s [3]int
	}{
		s : a,
	}

	//结构是值拷贝，内部的数组也是值拷贝
	d := c

	//修改c中的数组元素值并不影响a
	c.s[2] = 30

	//修改d中的数组元素值并不影响c
	d.s[2] = 20

	fmt.Printf("%p,%v\n",&a,a)
	fmt.Printf("%p,%v\n",&c,c)
	fmt.Printf("%p,%v\n",&d,d)
}
