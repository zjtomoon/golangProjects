package main

import (
	"fmt"
	"time")


func main() {
	//timer实现延时的功能
	time.Sleep(time.Second)
	//fmt.Println("1")
	timer3 := time.NewTimer(2*time.Second)
	<- timer3.C
	fmt.Println("2秒到")
	<-time.After(2*time.Second)
	fmt.Println("2秒到")
}