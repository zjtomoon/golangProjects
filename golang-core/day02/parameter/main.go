package main

import (
	"fmt"
)
func chvalue(a int) int{
	a = a + 1 
	return a
}
func chpointer(a *int){
	*a = *a + 1
	return
}
func main()  {
	a := 10
	chvalue(a) //实参传递给形参是值拷贝
	fmt.Println(a)

	chpointer(&a) //实参传递给形参仍然是值拷贝，只不过复制的是a的地址值
	fmt.Println(a)
}
