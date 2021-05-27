package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func write()  {
	rwlock.Lock()
	x += 1
	time.Sleep(10 * time.Microsecond) //假设写操作耗时10毫秒
	rwlock.Unlock()
	wg.Done()
}

func read()  {
	rwlock.RLock()
	time.Sleep(time.Microsecond) //假设读操作耗时1毫秒
	rwlock.RUnlock()
	wg.Done()

}

func main()  {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

//需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。