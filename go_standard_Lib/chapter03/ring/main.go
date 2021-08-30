package main

import (
	"container/ring"
	"fmt"
)

func main() {
	ring1 := ring.New(3)

	for i := 1; i <= 3; i++ {
		ring1.Value = i
		ring1 = ring1.Next()
	}

	//计算1+2+3
	s := 0
	ring1.Do(func(p interface{}) {
		s += p.(int)
	})

	fmt.Println("sum is ", s)
}

//sum is  6
