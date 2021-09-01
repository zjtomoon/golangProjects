package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	go func() {
		fmt.Println("hello world !")
		mu.Lock()
	}()

	mu.Unlock()
}

/*
并发编程的核心概念是同步通信，但是同步的方式却有多种。
我们先以大家熟悉的互斥量sync.Mutex来实现同步通信。
根据文档，我们不能直接对一个未加锁状态的sync.Mutex进行解锁，这会导致运行时异常。
下面这种方式并不能保证正常工作：
 */