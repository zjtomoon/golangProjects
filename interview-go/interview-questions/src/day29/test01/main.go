package main

import "fmt"

func main() {
	v := []int{1,2,3}
/*	for k,_ := range v {
		v = append(v,k)
	}*/

	for _, val := range v {
		v = append(v,val)
	}
	fmt.Println(v)
}
