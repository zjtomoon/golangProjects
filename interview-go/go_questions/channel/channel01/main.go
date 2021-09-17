package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 19
	close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received:", x)
	}
	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed,data invalid.")
	}
}
