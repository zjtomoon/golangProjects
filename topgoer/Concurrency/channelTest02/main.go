package main

import "fmt"

//有缓冲的通道

func recv(c chan int)  {
	res := <- c
	fmt.Println("接收成功,res:",res)
}

func main()  {
	ch := make(chan int,1)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功,ch:",ch)
}