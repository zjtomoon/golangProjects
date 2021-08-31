package main

import (
	"fmt"
	"sync"
)

func main() {
	wp := new(sync.WaitGroup)
	wp.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("done", i)
			wp.Done()
		}()
	}

	wp.Wait()
	fmt.Println("wait end")
}
