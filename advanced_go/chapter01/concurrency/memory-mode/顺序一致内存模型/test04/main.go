package main

import "fmt"

func test(c chan int)  {
	data := <- c
	fmt.Println(data)
}

func main() {
	c := make(chan int)
	go test(c)
	c <- 12
}
