package main

import (
	"fmt"
)

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	//错误写法
/*	for key,val := range slice {
		m[key] = &val
	}

	for k,v := range m {
		fmt.Println(k,"->",*v)
	}*/

	fmt.Println("=============")

	for key1,val1 := range slice {
		value := val1
		m[key1] = &value
	}

	for k1,v1 := range m {
		fmt.Println(k1,"====>",*v1)
	}
}
