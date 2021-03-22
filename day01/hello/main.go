package main

import (
	"fmt"
)
func main(){
	fmt.Println("hello world !")
	var a int = 1
	fmt.Println(a)
	b := 1
	fmt.Println(b)
	fmt.Println("====================")
	c:= "hello"
	d := "world"
	e := c + d
	for i := 0 ; i < len(e) ; i++ {
		fmt.Println(e[i])
	}
	fmt.Println("====================")
	for i,v := range e {
		fmt.Println(i,v)
	}
}