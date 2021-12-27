package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
				//fmt.Println("ch = 0")
			case ch <- 1:
			//fmt.Println("ch = 1")
			default:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
