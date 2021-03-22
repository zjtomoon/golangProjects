package main

import "fmt"

func main() {
	//使用字面量创建
	ma := map[string]int{"a": 1, "b": 2}
	fmt.Println(ma["a"])
	fmt.Println(ma["b"])

	//使用内置的make函数创建
	//make(map[K]T)
	//make(map[K]T,len)

	mp1 := make(map[int]string)
	mp2 := make(map[int]string, 10)
	mp1[1] = "tom"
	mp2[1] = "pony"
	fmt.Println(mp1[1]) //tom
	fmt.Println(mp2[1]) //pony
}
