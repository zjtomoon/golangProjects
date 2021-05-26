package main

import (
	"fmt"
	"time")


func main() {
	//获取ticker对象
	//时间到了，多次执行
	ticker := time.NewTicker(1*time.Second)
	i := 0
	//子协程
	go func()  {
		for{
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for{
		
	}
}