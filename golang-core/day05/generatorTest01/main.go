package main

//最简单的带缓冲的生成器

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	//启动一个goroutinue用于生产随机数，函数返回一个通道用于获取随机数
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main() {
	ch := GenerateIntA()
	for i := 0 ; i < 10 ; i++{
		fmt.Println(<- ch)
	}
}
