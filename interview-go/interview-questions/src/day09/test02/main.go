package main

import "fmt"

func hello(num ...int)  {
	num[0] = 19
}

func main() {
	i := []int{5,6,7}
	hello(i...)
	fmt.Println(i[0])
}
