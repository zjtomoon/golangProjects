package main

import "fmt"

func main() {
	done := make(chan int,10) //带10个缓存

	//开启N个后台打印线程

	for i := 0 ; i < cap(done) ; i++ {
		go func() {
			fmt.Println("hello world !")
			done <- 1
		}()
	}

	//等待N个后台线程完成
	for i := 0 ; i <cap(done) ; i++ {
		<- done
	}
}
