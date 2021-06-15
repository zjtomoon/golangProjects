package main

import (
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	time.Sleep(1 * time.Second)
}

func main() {
	go sum()
	//NumGoroutinue返回当前程序的goroutinue数目
	println("NumGoRoutinue=", runtime.NumGoroutine())

	//main goroutinue故意“sleep” 5s,防止其提前退出

	time.Sleep(5 * time.Second)
}
