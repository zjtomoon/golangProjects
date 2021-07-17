package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func Printer(ch chan int)  {
	for i := range ch {
		fmt.Printf("Received :%d\n",i)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	go Printer(c)
	wg.Add(1)
	for i := 1; i <= 10 ; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
