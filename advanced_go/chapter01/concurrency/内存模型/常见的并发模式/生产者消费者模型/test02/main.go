package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//生产者：生成factor 整数倍的序列

func Producer(factor int,out chan<- int)  {
	for i := 0 ; ; i++ {
		out <- i*factor
	}
}

//消费者

func Consumer(in <-chan int)  {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int,64) //成果队列

	go Producer(3,ch)
	go Producer(5,ch)
	go Consumer(ch)

	//Ctrl+C退出
	sig := make(chan os.Signal,1)
	signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
	fmt.Printf("quit (%v)\n",<-sig)
}
