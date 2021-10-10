package main

import "fmt"

func main() {
	s1 := []int{1,2,3}
	s2 := []int{4,5,6}
	arr := append(s1,s2...)
	fmt.Printf("arr类型为%T\n",arr)
	fmt.Println(arr)
}
