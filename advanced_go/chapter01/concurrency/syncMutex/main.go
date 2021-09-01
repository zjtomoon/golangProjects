package main

import (
	"fmt"
	"sync"
)

/*
	一般情况下，原子操作都是通过“互斥”访问来保证的，通常由特殊的cpu指令提供保护。
	当然，仅仅是模拟粗粒度的原子操作，可以借助于sync.Mutex来实现。
 */

var total struct{
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 0 ; i <= 100 ; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}
