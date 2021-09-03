package main

import "fmt"

func main() {
	done := make(chan int)

	go func() {
		fmt.Println("hello world !")
		<- done
	}()

	done <- 1
}
