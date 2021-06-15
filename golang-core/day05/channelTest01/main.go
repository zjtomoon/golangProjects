package main

import "runtime"

func main() {
	c := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		//写通道
		c <- struct{}{}
	}(c)
	println("NumGORoutinue=", runtime.NumGoroutine())
	//读通道c 通过通道进行同步等待
	<-c

}
