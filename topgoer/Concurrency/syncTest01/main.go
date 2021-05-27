package main

import (
	"fmt"
	"sync")


var wg sync.WaitGroup

func hello()  {
	defer wg.Done()
	fmt.Println("Hello Goroutinue!")
}

func main() {
	wg.Add(1)
	go hello() //启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutinue done!")
	wg.Wait()
}