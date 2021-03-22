package main

import (
	"fmt"
)

func main()  {
	a := [3]int{1,2,3}
	//a0 := a[0]
	for i,v := range a {
		fmt.Println(a[i],i,v)
	}
	fmt.Println("==========")
	b := [...]int{1,2,3}
	//blength = len(b)
	for i := 0 ; i < len(b) ; i++ {
		fmt.Println(b[i])
	}
}