package main

import (
	"fmt"
	"time"
)

//我们通过select和default分支可以很容易实现一个Goroutine的退出控制

func worker(channel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			//正常工作
		case <-channel:
			//退出
		}
	}
}

func main() {
	channel := make(chan bool)
	go worker(channel)

	time.Sleep(time.Second)
	channel <- true //提供一个停止的信号
}
