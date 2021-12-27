package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
			return
		}
	}
}

func main() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	//wg.Wait()
}

/*
	我们通过close来关闭cancel管道向多个goroutine广播退出的指令。不过这个程序依然不够稳健：
	当每个Goroutine收到退出指令退出时会进行一定的清理工作，但是退出的清理工作并不能保证被完成，
	因为main线程并没有等待各个工作goroutine退出工作完成的机制。故可以结合sync.WaitGroup来改进
*/
