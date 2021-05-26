package main

import (
	"fmt"
	"time")


func main() {
	//验证timer只能响应一次
	timer2 := time.NewTimer(2*time.Second)
	for{
		<- timer2.C
		fmt.Println("时间到")
	}
}