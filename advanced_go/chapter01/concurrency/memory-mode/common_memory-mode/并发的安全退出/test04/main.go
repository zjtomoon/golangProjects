package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup,channel chan bool)  {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <- channel:
			return
		}
	}
}

func main() {
	channel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go worker(&wg,channel)
	}

	time.Sleep(time.Second)
	close(channel)
	wg.Wait()
}
