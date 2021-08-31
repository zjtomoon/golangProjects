package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.Cond是用来控制某个条件下，goroutine进入等待时期，等待信号到来，
//然后重新启动
func main() {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	done := false

	cond.L.Lock()

	go func() {
		time.Sleep(2e9)
		done = true
		cond.Signal()
	}()

	if !done {
		cond.Wait()
	}
	fmt.Println("new done is ", done)
}

//这里当主goroutine进入cond.Wait的时候，就会进入等待，当从goroutine发出信号之后，
// 主goroutine才会继续往下面走。
