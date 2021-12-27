package main

import (
	"fmt"
	"time"
)

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

	for i := 0; i < 10; i++ {
		go worker(channel)
	}
	time.Sleep(time.Second)
	close(channel)
	//管道的发送操作和接收操作是一一对应的，如果要停止多个goroutine可能需要创建相同数量的管道，代价太大。
	//可以通过close关闭一个管道来实现广播的效果，所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标记。
}
