package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功,ret:",ret)
}

func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}