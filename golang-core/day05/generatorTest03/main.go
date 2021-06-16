package main

import (
	"fmt"
	"math/rand"
)

//有时希望生产器能够自动退出，可以借助Go通道的退出机制

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			//通过select 监听一个信号chan 来确定是否停止生成
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//不再需要生成器，通过close chan发送一个通知给生成器
	close(done)
	for v := range ch {
		fmt.Println(v)
	}
}
