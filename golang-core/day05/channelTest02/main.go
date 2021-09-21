package main

import "runtime"

/*
	goroutinue运行结束后退出，写到缓冲通道中的数据不会消失，它可以缓冲和适配两个goroutinue处理速率不一致的情况，缓冲通道和消息队列
	类似，有削峰和增大吞吐量的功能
*/
func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		c <- struct{}{}
	}(c, ci)

	//NumGoroutine可以返回当前程序的goroutine数目
	println("NmGoroutinue=", runtime.NumGoroutine())

	//读通道c，通过通道进行同步等待
	<- c
	//此时ci通道已经关闭，匿名函数启动的goroutine已经退出
	println("NumGrotinue=", runtime.NumGoroutine())

	//但通道ci还可以继续读取
	for v := range ci {
		println(v)
	}
}
