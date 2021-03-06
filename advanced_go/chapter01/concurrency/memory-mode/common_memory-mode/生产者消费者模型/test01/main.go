package main

import (
	"fmt"
	"time"
)

/*
并发编程中最常见的例子就是生产者消费者模式，
该模式主要通过平衡生产线程和消费线程的工作能力来提高程序的整体处理数据的速度。
简单地说，就是生产者生产一些数据，然后放到成果队列中，同时消费者从成果队列中来取这些数据。
这样就让生产消费变成了异步的两个过程。当成果队列中没有数据时，消费者就进入饥饿的等待中；
而当成果队列中数据已满时，生产者则面临因产品挤压导致CPU被剥夺的下岗问题。
 */
 
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

	go Producer(3,ch) //生成 3 的倍数的序列
	go Producer(5,ch) //生成 5 的倍数的序列
	go Consumer(ch) //消费 生成的队列

	//运行一定时间后退出
	time.Sleep(5 * time.Second)
}
