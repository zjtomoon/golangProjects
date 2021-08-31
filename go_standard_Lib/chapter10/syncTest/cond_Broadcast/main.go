package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() //获取锁
	cond.Wait()   //等待通知 暂时阻塞
	fmt.Println(x)
	time.Sleep(1 * time.Second)
	cond.L.Unlock() //释放锁，不释放的话将只会有一次输出
}

func main() {
	for i := 0; i < 40; i++ {
		go test(i)
	}
	fmt.Println("start all")
	cond.Broadcast() //下发广播给所有等待的goroutinue
	time.Sleep(60 * time.Second)
}
