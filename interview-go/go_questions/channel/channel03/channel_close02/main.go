package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const numSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	//senders
	for i := 0; i < numSenders; i++ {
		go func() {
			select {
			case <-stopCh:
				return
			case dataCh <- rand.Intn(Max):
			}
		}()
	}

	//the receivers
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop singnal to senders.")
				close(stopCh)
				return
			}

			fmt.Println(value)
		}
	}()

	select {
	case <-time.After(time.Hour):
	}

}
