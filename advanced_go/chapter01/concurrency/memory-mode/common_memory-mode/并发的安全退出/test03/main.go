package main

import (
	"fmt"
	"time"
)

func worker(channel chan bool)  {
	for {
		select {
		default:
			fmt.Println("hello")
			//正常工作
		case <- channel:
			//退出
		}
	}
}

func main() {
	channel := make(chan bool)

	for i := 0 ; i < 10 ; i++ {
		go worker(channel)
	}
	 time.Sleep(time.Second)
	close(channel)
}
