package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	a := add(1,2)
	b := add(1,3,7)
	c := add([]int{1,3,7}...)

	fmt.Printf("a = %d , b = %d, c = %d\n",a,b,c)
}
