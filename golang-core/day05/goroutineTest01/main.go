package main

import (
	"runtime"
	"time"
)

func main() {
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		time.Sleep(1 * time.Second)
	}()

	//NumGoroutinue可以返回当前程序的goroutinue数目
	println("NumGoroutinue=", runtime.NumGoroutine())

	//main goroutinue故意"sleep" 5秒，防止其提前推出
	time.Sleep(5 * time.Second)
}
