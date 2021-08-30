package main

import (
	"time"
	"fmt"
)

//通过time.After模拟超时

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		<- c
	}()

	select {
	case c <- 1:
		fmt.Println("channel...")
	case <- time.After(2 * time.Second):
		close(c)
		fmt.Println("timeout...")
	}
	start := time.Now()
	timer := time.AfterFunc(2 * time.Second, func() {
		fmt.Println("after func callback,elaspe:",time.Now().Sub(start))
	})
	time.Sleep(1 * time.Second)
	//time.Sleep(3 * time.Second)
	//time.Stop停止定时器或time.Reset重置定时器
	//Reset 在Timer还未触发时返回true;触发了或stop了，返回false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger !")
	} else {
		fmt.Println("timer had expired or stop!")
	}
	time.Sleep(10 * time.Second)
}
