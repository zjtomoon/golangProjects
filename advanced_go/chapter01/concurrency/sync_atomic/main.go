package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	用互斥锁来保护一个数值型的共享资源麻烦，且效率低下。标准库的sync/atomic包对原子操作提供了丰富的支持。
 */

var total uint64

func worker(wg *sync.WaitGroup)  {
	defer wg.Done()

	var i uint64
	for i = 0 ; i <= 100 ; i++ {
		atomic.AddUint64(&total,i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total)
}
