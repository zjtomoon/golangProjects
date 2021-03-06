package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//开启N个后台打印线程
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("hello world !")
			wg.Add(-1)
			//wg.Done()
		}()
	}

	//等待N个后台线程完成
	wg.Wait()
}
